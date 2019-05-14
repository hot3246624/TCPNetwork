package client

import (
	"github.com/cosmos/cosmos-sdk/client"

	tcpcmd "github.com/hot3246624/TCPNetwork/x/tcp/client/cli"

	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	tcpTxCmd := &cobra.Command{
		Use:   "tcp",
		Short: "TCPNetwork transactions subcommands",
	}

	tcpTxCmd.AddCommand(client.PostCommands(
		tcpcmd.GetCmdTransfer(mc.cdc),
	)...)

	return tcpTxCmd
}
