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
		Use:     "register-admin [name]",
		Short:   "registers an admin",
		Long:    `registers an admin`,
		Example: "./simd tx lms register-admin adminname --from adminaddress --chain-id testnet",
		Args:    cobra.ExactArgs(1),
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
		Use:     "add-students [studentaddress] [studentname] [studentid] ...",
		Short:   "add students",
		Long:    `add students`,
		Example: "./simd tx lms add-students cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u saiteja 2045 --from adminaddress --chain-id testnet",
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
		Use:     "apply-leave [reason] [from] [to]",
		Short:   "to apply a leave",
		Long:    `to apply a leave`,
		Example: "./simd tx lms apply-leave sick 12-Jan-2022 13-Jan-2022 --from studentaddress --chain-id testnet",
		Args:    cobra.ExactArgs(3),
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
				Leave: &types.Leave{
					Address:   fromAddress,
					Reason:    args[0],
					From:      &from,
					To:        &to,
					HandledBy: "no one Hanlded it yet",
					Status:    types.LeaveStatus_STATUS_PENDING,
				},
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
		Use:     "accept-leave [student address] [y/n]",
		Short:   "For admin to Accept a leave",
		Long:    `For admin to Accept a leave`,
		Example: "./simd tx lms accept-leave cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u --from adminaddress --chain-id testnet",
		Args:    cobra.ExactArgs(2),
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
