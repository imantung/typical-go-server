package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typicli"
	"github.com/typical-go/typical-go/pkg/typiobj"
	"github.com/typical-go/typical-rest-server/app"
)

func TestModule(t *testing.T) {
	m := app.Module()
	require.True(t, typiobj.IsProvider(m))
	require.True(t, typiobj.IsConfigurer(m))
	require.True(t, typiobj.IsPreparer(m))
	require.True(t, typiobj.IsRunner(m))
	require.True(t, typicli.IsAppCommander(m))
}
