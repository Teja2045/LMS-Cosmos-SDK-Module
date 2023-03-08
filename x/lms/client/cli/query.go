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
		GetCmdListPendingLeaves(),
		GetCmdListStudents(),
		GetCmdListHandledLeaves(),
		GetCmdListAllAcceptedLeaves(),
		GetCmdListAllRejectedLeaves(),
	)
	return lmsQueryCmd
}

// GetCmdcheckStatus return status of a leave for a student.
func GetCmdLeavetatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "leave-status [student address] [student name]",
		Short:   "retrieves if the last applied leave's status",
		Long:    `retrieves if the last applied leave's status`,
		Args:    cobra.ExactArgs(2),
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

// GetCmdListPendingLeaves returns list of leaves that are still pending.
func GetCmdListPendingLeaves() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-pending-leaves [admin adress] [admin name]",
		Short:   "list out the pending leaves",
		Args:    cobra.ExactArgs(2),
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
			res, err := queryClient.ListPendingLeaves(context.Background(), listLeavesRequest)
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
		Args:    cobra.ExactArgs(1),
		Example: "./simd query lms list-students cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu",
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

// GetCmdListHandledLeaves returns list of leaves that are handled by an admin.
func GetCmdListHandledLeaves() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-handled-leaves [admin adress]",
		Short:   "list out the handled leaves by that admin",
		Args:    cobra.ExactArgs(1),
		Example: "./simd query lms list-handled-leaves cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu adminname",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			listHandledLeavesRequest := &types.ListHandledLeavesRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListHandledLeaves(context.Background(), listHandledLeavesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdListAllRejectedLeaves returns list of leaves that are rejected.
func GetCmdListAllRejectedLeaves() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-rejected-leaves [admin adress]",
		Short:   "list out the handled leaves by that admin",
		Example: "./simd query lms list-handled-leaves cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu adminname",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			listHandledLeavesRequest := &types.ListHandledLeavesRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListAllRejectedLeaves(context.Background(), listHandledLeavesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdAllAcceptedLeaves returns list of leaves that are accepted.
func GetCmdListAllAcceptedLeaves() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-accepted-leaves [admin adress]",
		Short:   "list out the handled leaves by that admin",
		Args:    cobra.ExactArgs(1),
		Example: "./simd query lms list-handled-leaves cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu adminname",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			listHandledLeavesRequest := &types.ListHandledLeavesRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListAllAcceptedLeaves(context.Background(), listHandledLeavesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
