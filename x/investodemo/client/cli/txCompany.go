package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/faddat/investo-demo/x/investodemo/types"
)

func CmdCreateCompany() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-company [mktgroi] [logo] [netflow] [industry]",
		Short: "Creates a new company",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsMktgroi := string(args[0])
			argsLogo := string(args[1])
			argsNetflow := string(args[2])
			argsIndustry := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCompany(clientCtx.GetFromAddress().String(), string(argsMktgroi), string(argsLogo), string(argsNetflow), string(argsIndustry))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCompany() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-company [id] [mktgroi] [logo] [netflow] [industry]",
		Short: "Update a company",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsMktgroi := string(args[1])
			argsLogo := string(args[2])
			argsNetflow := string(args[3])
			argsIndustry := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCompany(clientCtx.GetFromAddress().String(), id, string(argsMktgroi), string(argsLogo), string(argsNetflow), string(argsIndustry))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteCompany() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-company [id] [mktgroi] [logo] [netflow] [industry]",
		Short: "Delete a company by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCompany(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
