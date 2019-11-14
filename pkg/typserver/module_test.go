package typserver_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typiobj"
	"github.com/typical-go/typical-rest-server/pkg/typserver"
)

func TestModule(t *testing.T) {
	m := typserver.Module()
	require.True(t, typiobj.IsProvider(m))
	require.True(t, typiobj.IsDestroyer(m))
	require.True(t, typiobj.IsConfigurer(m))
}
