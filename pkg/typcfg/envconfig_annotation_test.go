package typcfg_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-rest-server/pkg/typcfg"
)

func TestCfgAnnotation_Annotate(t *testing.T) {
	os.MkdirAll("somepkg1", 0777)
	defer os.RemoveAll("somepkg1")

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	var out strings.Builder
	typcfg.Stdout = &out
	defer func() { typcfg.Stdout = os.Stdout }()

	EnvconfigAnnotation := &typcfg.EnvconfigAnnotation{}
	c := &typast.Context{
		Destination: "somepkg1",
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
		Summary: &typast.Summary{
			Annots: []*typast.Annot{
				{
					TagName: "@envconfig",
					Decl: &typast.Decl{
						File: typast.File{Package: "mypkg"},
						Type: &typast.StructDecl{
							TypeDecl: typast.TypeDecl{Name: "SomeSample"},
							Fields: []*typast.Field{
								{Names: []string{"SomeField1"}, Type: "string", StructTag: `default:"some-text"`},
								{Names: []string{"SomeField2"}, Type: "int", StructTag: `default:"9876"`},
							},
						},
					},
				},
			},
		},
	}

	require.NoError(t, EnvconfigAnnotation.Annotate(c))

	b, _ := ioutil.ReadFile("somepkg1/envconfig_annotated.go")
	require.Equal(t, `package somepkg1

/* Autogenerated by Typical-Go. DO NOT EDIT.

TagName:
	@envconfig

Help:
	https://pkg.go.dev/github.com/typical-go/typical-rest-server/pkg/typcfg
*/

import (
	"github.com/kelseyhightower/envconfig"
)

func init() { 
	typapp.AppendCtor(
		&typapp.Constructor{Name: "",Fn: LoadSomeSample},
	)
}

// LoadSomeSample load env to new instance of SomeSample
func LoadSomeSample() (*mypkg.SomeSample, error) {
	var cfg mypkg.SomeSample
	if err := envconfig.Process("SOMESAMPLE", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
`, string(b))

	require.Equal(t, "Generate @envconfig to somepkg1/envconfig_annotated.go\n", out.String())

}

func TestCfgAnnotation_Annotate_GenerateDotEnvAndUsageDoc(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	var out strings.Builder
	typcfg.Stdout = &out
	defer func() { typcfg.Stdout = os.Stdout }()

	defer os.Clearenv()

	a := &typcfg.EnvconfigAnnotation{
		Target:   "some-target",
		Template: "some-template",
		DotEnv:   ".env33",
		UsageDoc: "some-usage.md",
	}
	c := &typast.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
		Summary: &typast.Summary{Annots: []*typast.Annot{
			{
				TagName:  "@envconfig",
				TagParam: `ctor_name:"ctor1" prefix:"SS"`,
				Decl: &typast.Decl{
					File: typast.File{Package: "mypkg"},
					Type: &typast.StructDecl{
						TypeDecl: typast.TypeDecl{Name: "SomeSample"},
						Fields: []*typast.Field{
							{Names: []string{"SomeField1"}, Type: "string", StructTag: `default:"some-text"`},
							{Names: []string{"SomeField2"}, Type: "int", StructTag: `default:"9876"`},
						},
					},
				},
			},
		}},
	}

	require.NoError(t, a.Annotate(c))
	defer os.Remove(a.Target)
	defer os.Remove(a.DotEnv)
	defer os.Remove(a.UsageDoc)

	b, _ := ioutil.ReadFile(a.Target)
	require.Equal(t, `some-template`, string(b))

	b, _ = ioutil.ReadFile(a.DotEnv)
	require.Equal(t, "SS_SOMEFIELD1=some-text\nSS_SOMEFIELD2=9876\n", string(b))
	require.Equal(t, "some-text", os.Getenv("SS_SOMEFIELD1"))
	require.Equal(t, "9876", os.Getenv("SS_SOMEFIELD2"))

	require.Equal(t, "Generate @envconfig to some-target\nNew keys added in '.env33': SS_SOMEFIELD1 SS_SOMEFIELD2\nGenerate 'some-usage.md'\n", out.String())
}

func TestCfgAnnotation_Annotate_Predefined(t *testing.T) {
	target := "cfg-target"

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)
	defer os.RemoveAll(target)

	EnvconfigAnnotation := &typcfg.EnvconfigAnnotation{
		TagName:  "@some-tag",
		Template: "some-template",
		Target:   target,
	}
	c := &typast.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
		Summary: &typast.Summary{
			Annots: []*typast.Annot{
				{
					TagName: "@some-tag",
					Decl: &typast.Decl{
						File: typast.File{Package: "mypkg"},
						Type: &typast.StructDecl{
							TypeDecl: typast.TypeDecl{Name: "SomeSample"},
							Fields:   []*typast.Field{},
						},
					},
				},
			},
		},
	}
	require.NoError(t, EnvconfigAnnotation.Annotate(c))

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, `some-template`, string(b))
}

func TestCfgAnnotation_Annotate_RemoveTargetWhenNoAnnotation(t *testing.T) {
	target := "target1"
	defer os.Remove(target)
	ioutil.WriteFile(target, []byte("some-content"), 0777)
	c := &typast.Context{
		Context: &typgo.Context{},
		Summary: &typast.Summary{},
	}

	EnvconfigAnnotation := &typcfg.EnvconfigAnnotation{Target: target}
	require.NoError(t, EnvconfigAnnotation.Annotate(c))
	_, err := os.Stat(target)
	require.True(t, os.IsNotExist(err))
}

func TestCreateField(t *testing.T) {
	testnames := []struct {
		TestName string
		Prefix   string
		Field    *typast.Field
		Expected *typcfg.Field
	}{
		{
			Prefix:   "APP",
			Field:    &typast.Field{Names: []string{"Address"}},
			Expected: &typcfg.Field{Key: "APP_ADDRESS"},
		},
		{
			Prefix: "APP",
			Field: &typast.Field{
				Names:     []string{"some-name"},
				StructTag: reflect.StructTag(`envconfig:"ADDRESS" default:"some-address" required:"true"`),
			},
			Expected: &typcfg.Field{Key: "APP_ADDRESS", Default: "some-address", Required: true},
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typcfg.CreateField(tt.Prefix, tt.Field))
		})
	}
}