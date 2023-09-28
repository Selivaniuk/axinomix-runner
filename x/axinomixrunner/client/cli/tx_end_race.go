package cli

import (
	"strconv"

	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEndRace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end-race [id] [coins] [score]",
		Short: "Broadcast message end-race",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argCoins, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argScore, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEndRace(
				clientCtx.GetFromAddress().String(),
				argId,
				argCoins,
				argScore,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
