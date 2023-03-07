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
		GetCmdListStudents(),
	)
	return lmsQueryCmd
}

// GetCmdcheckStatus return status of a leave for a student.
func GetCmdLeavetatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "leave-status [student address] [student name]",
		Short:   "retrieves if the last applied leave's status",
		Long:    `retrieves if the last applied leave's status`,
		Example: "./simd query lms leave-status cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u studentname",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			leaveStatusRequest := &types.LeaveStatusRequest{
				Address: args[0],
				Name:    args[1],
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.LeaveStatus(context.Background(), leaveStatusRequest)
			if err != nil {
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
		Use:     "list-pending-leaves [admin adress] [admin name]",
		Short:   "list out the pending leaves",
		Example: "./simd query lms list-pending-leaves cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu adminname",
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

func GetCmdListStudents() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-students",
		Short:   "listout all the students",
		Long:    `listout all the students for admin`,
		Example: "./simd query lms list-students cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			listStudentRequest := &types.ListStudentsRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListStudents(context.Background(), listStudentRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
