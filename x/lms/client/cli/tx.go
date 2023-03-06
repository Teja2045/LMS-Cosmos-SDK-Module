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
		Use:   "RegisterAdmin [address] [name]",
		Short: "| Name |",
		Long:  `registers admin`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()
			admin := types.MsgRegisterAdminRequest{
				Address:       fromAddress,
				Name:          args[0],
				SignerAddress: fromAddress,
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
		Short: " | student1 | student2 |.... (student{address,name,id})",
		Long:  `registers admin`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()

			students := []*types.Student{}
			for i := 0; i < len(args)/3; i++ {
				student := &types.Student{
					Address: args[3*i],
					Name:    args[3*i+1],
					Id:      args[3*i+2],
				}
				students = append(students, student)
			}
			AddStudents := types.MsgAddStudentRequest{
				Admin:         fromAddress,
				Students:      students,
				SignerAddress: fromAddress,
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
		Short: "| reason | from | to |",
		Long:  `to apply leave`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()
			var format string = "2006-Jan-06"
			from, _ := time.Parse(format, args[1])
			to, _ := time.Parse(format, args[2])
			applyleave := types.MsgApplyLeaveRequest{
				Address:       fromAddress,
				Reason:        args[0],
				From:          &from,
				To:            &to,
				SignerAddress: fromAddress,
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
		Short: "|student address| y(for accepted, anything else rejected) |",
		Long:  `For admin to Accept a leave`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddress := clientCtx.GetFromAddress().String()
			status := types.LeaveStatus_STATUS_REJECTED
			if args[1] == "y" {
				status = types.LeaveStatus_STATUS_ACCEPTED
			}

			admin := types.MsgAcceptLeaveRequest{
				Admin:         fromAddress,
				Student:       args[0],
				Status:        status,
				SignerAddress: fromAddress,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &admin)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
