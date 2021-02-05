package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/faddat/investo-demo/x/investodemo/types"
	"github.com/spf13/cobra"
)

func CmdListCompany() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-company",
		Short: "list all company",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCompanyRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CompanyAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowCompany() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-company [id]",
		Short: "shows a company",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCompanyRequest{
				Id: args[0],
			}

			res, err := queryClient.Company(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
