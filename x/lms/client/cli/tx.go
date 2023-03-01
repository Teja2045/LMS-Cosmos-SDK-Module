package cli

import (
	"lmsmodule/x/lms/types"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	nftTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "|lms|",
		Long:  `lms module commands`,
		RunE:  client.ValidateCmd,
	}

	nftTxCmd.AddCommand(
		NewCmdRegisterAdmin(),
		NewCmdAddStudents(),
		NewCmdAcceptLeave(),
		NewCmdApplyLeave(),
	)

	return nftTxCmd
}

// NewCmdRegisterAdmin creates a CLI command for MsgRegisterAdminRequest.
func NewCmdRegisterAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "RegisterAdmin",
		Short: "| address | Name |",
		Long:  `registers admin`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			admin := types.MsgRegisterAdminRequest{
				Address: args[0],
				Name:    args[1],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &admin)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCmdAddStudents creates a CLI command for MsgAddStudents.
func NewCmdAddStudents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "AddStudents",
		Short: " | adminAddress | student1 | student2 |.... (student{address,name,id})",
		Long:  `registers admin`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			adminAddress := args[0]
			students := []*types.Student{}
			for i := 0; i < (len(args)-1)/3; i++ {
				student := &types.Student{
					Address: args[3*i+1],
					Name:    args[3*i+2],
					Id:      args[3*i+3],
				}
				students = append(students, student)
			}
			AddStudents := types.MsgAddStudentRequest{
				Admin:    adminAddress,
				Students: students,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &AddStudents)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCmdApplyLeave creates a CLI command for MsgApplyLeave.
func NewCmdApplyLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ApplyLeave",
		Short: "| address | reason | from | to |",
		Long:  `to apply leave`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			var format string = "2006-Jan-06"
			from, _ := time.Parse(format, args[2])
			to, _ := time.Parse(format, args[3])
			applyleave := types.MsgApplyLeaveRequest{
				Address: args[0],
				Reason:  args[1],
				From:    &from,
				To:      &to,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &applyleave)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCmdAcceptLeave creates a CLI command for MsgAcceptLeave.
func NewCmdAcceptLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "AcceptLeave",
		Short: "| admin address | student address |",
		Long:  `For admin to Accept a leave`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			admin := types.MsgAcceptLeaveRequest{
				Admin:   args[0],
				Student: args[1],
				Status:  types.LeaveStatus_STATUS_ACCEPTED,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &admin)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
