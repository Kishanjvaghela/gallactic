package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoading(t *testing.T) {
	conf1 := DefaultConfig()
	conf1.Tendermint.ListenAddress = "1.1.1.1"
	conf1.Tendermint.Moniker = "moniker-test"
	conf1.Tendermint.TendermintRoot = "tendermint1"
	conf1.GRPC.Enabled = false
	conf1.RPC.Enabled = true
	toml, err := conf1.ToTOML()
	require.NoError(t, err)
	fmt.Println(toml)
	conf2, err := FromTOML(string(toml))
	require.NoError(t, err)
	require.Equal(t, conf1, conf2)
}
