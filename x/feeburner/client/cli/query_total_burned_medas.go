package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/oxygene76/medasdigital/x/feeburner/types"
)

func CmdQueryTotalBurnedMedasAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-burned-medas-amount",
		Short: "shows total amount of burned medas",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalBurnedMedasAmount(context.Background(), &types.QueryTotalBurnedMedasAmountRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
