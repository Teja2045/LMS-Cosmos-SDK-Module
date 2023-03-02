package cli

import (
	"context"

	"github.com/spf13/cobra"

	"lmsmodule/x/lms/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	lmsQueryCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the lms module",
		Long:  ``,
		RunE:  client.ValidateCmd,
	}

	lmsQueryCmd.AddCommand(
		GetCmdLeavetatus(),
		GetCmdListLeaves(),
	)
	return lmsQueryCmd
}

// GetCmdcheckStatus return status of a leave for a student.
func GetCmdLeavetatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "LeaveStatus",
		Short: "| student Address | student name |",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			leaveStatusRequest := &types.LeaveStatusRequest{
				Address: args[0],
				Name:    args[1],
			}

			//fmt.Println("args", args)
			//fmt.Println(leaveStatusRequest)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.LeaveStatus(context.Background(), leaveStatusRequest)
			//fmt.Println("err", err)
			if err != nil {
				//fmt.Println("here in client", "\n", "error: ", err, "\n", "res: ", res)
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdListLeaves returns list of leaves that are still pending.
func GetCmdListLeaves() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ListLeaves",
		Short: "| admin Address | admin name |",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			listLeavesRequest := &types.ListLeavesRequest{
				Address: args[0],
				Name:    args[1],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListLeaves(context.Background(), listLeavesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
