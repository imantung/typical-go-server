package typical

import (
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/typical-go/typical-rest-server/pkg/typpg"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/server"
)

var (
	mainDB = typpg.Init(&typpg.Settings{
		DBName: "MyLibrary",
	})

	redis = typredis.Init(&typredis.Settings{})
)

// Descriptor of Typical REST Server
// Build-Tool and Application will be generated based on this descriptor
var Descriptor = typgo.Descriptor{

	Name:        "typical-rest-server",
	Description: "Example of typical and scalable RESTful API Server for Go",
	Version:     "0.8.29",

	EntryPoint: server.Main,

	Layouts: []string{"server", "pkg"},

	Configurer: typgo.Configurers{
		typredis.Configuration(redis),
		typpg.Configuration(mainDB),
		server.Configuration(),
	},

	Build: typgo.Builds{
		&typgo.StdBuild{},
		&typgo.Github{
			Owner:    "typical-go",
			RepoName: "typical-rest-server",
		},
	},

	Utility: typgo.Utilities{
		typpg.Utility(mainDB),   // create db, drop, migrate, seed, console, etc.
		typredis.Utility(redis), // redis console
		typmock.Utility(),

		typdocker.Compose(
			typpg.DockerRecipeV3(mainDB),
			typredis.DockerRecipeV3(redis),
		),
	},
}
