package cli

import (
	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLeaderboard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-leaderboard",
		Short: "list all leaderboard",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLeaderboardRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LeaderboardAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLeaderboard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-leaderboard [address]",
		Short: "shows a leaderboard",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			address := args[0]
			//id, err := strconv.ParseUint(args[0], 10, 64)
			//if err != nil {
			//	return err
			//}

			params := &types.QueryGetLeaderboardRequest{
				Address: address,
			}

			res, err := queryClient.Leaderboard(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
