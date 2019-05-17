package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/hot3246624/TCPNetwork/x/tcp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdTransfer is the CLI command for sending coins
func GetCmdTransfer(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer from to amount",
		Short: "transfer coins",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			from := args[0]
			to := args[1]
			amount := args[2]

			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			_, err := sdk.AccAddressFromBech32(from)
			if err != nil {
				return err
			}

			// get to address
			toAddr, err := sdk.AccAddressFromBech32(to)
			if err != nil {
				return err
			}

			// get transfer amount
			coins, err := sdk.ParseCoins(amount)
			if err != nil {
				return err
			}

			// TODO
			msg := tcp.NewMsgTransfer(cliCtx.GetFromAddress(), toAddr, coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	return cmd
}

// GetCmdContractDeploy is the CLI command for deploying contract
func GetCmdContractDeploy(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy [name] [value]",
		Short: "deploy contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			fromAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// TODO
			msg := tcp.NewMsgContractDeploy(fromAddr, []byte(args[1]))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}

// GetCmdContractExec is the CLI command for deploying contract
func GetCmdContractExec(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "exec [name] [value]",
		Short: "exec contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			fromAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := tcp.NewMsgContractExec(fromAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}
