package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"


	"github.com/hot3246624/TCPNetwork/x/tcp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

const (
	flagFrom   = "Fromx"
	flagTo     = "to"
	flagAmount = "amount"
)

// GetCmdTransfer is the CLI command for sending coins
func GetCmdTransfer(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer from to amount",
		Short: "transfer coins",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))


			to := viper.GetString(flagTo)
			amount := viper.GetString(flagAmount)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				fmt.Println("from account not exists")
				return err
			}

			// get from address
			fromAddr := cliCtx.GetFromAddress()

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
			msg := tcp.NewMsgTransfer(fromAddr, toAddr, coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagTo, "", "to address")
	cmd.Flags().String(flagAmount, "", "coin amount")
	cmd.MarkFlagRequired(flagTo)
	cmd.MarkFlagRequired(flagAmount)


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
