package types

import (
	//"lmsmodule/x/lms/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgAddStudentRequest{}
	_ sdk.Msg = &MsgRegisterAdminRequest{}
	_ sdk.Msg = &MsgApplyLeaveRequest{}
	_ sdk.Msg = &MsgAcceptLeaveRequest{}
)

func NewMsgAddStudentRequest(adminAddress string, students []*Student, signer string) *MsgAddStudentRequest {

	return &MsgAddStudentRequest{
		Admin:         adminAddress,
		Students:      students,
		SignerAddress: signer,
	}
}

func (msg MsgAddStudentRequest) ValidateBasic() error {
	eg := msg.Admin
	if _, err := sdk.AccAddressFromBech32(eg); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}

	if len(msg.Students) == 0 {
		return ErrNoNewStudentToBeAdded
	}
	return nil
}

func (msg MsgAddStudentRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgAddStudentRequest) GetSigners() []sdk.AccAddress {
	// sign := msg.Admin
	addr, _ := sdk.AccAddressFromBech32(msg.SignerAddress)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________

func NewMsgApplyLeaveRequest(address string, reason string, from *time.Time, to *time.Time, signer string) *MsgApplyLeaveRequest {

	return &MsgApplyLeaveRequest{
		Address:       address,
		Reason:        reason,
		From:          from,
		To:            to,
		SignerAddress: signer,
	}
}

func (msg MsgApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	if msg.Address == "" {
		return ErrAdminDetailsNil
	}
	if msg.Reason == "" {
		return ErrLeaveDetailsNil
	}
	if msg.From == nil {
		return ErrLeaveDetailsNil
	}
	if msg.To == nil {
		return ErrLeaveDetailsNil
	}
	return nil
}

func (msg MsgApplyLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	sign := msg.SignerAddress
	addr, _ := sdk.AccAddressFromBech32(sign)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________

func NewMsgRegisterAdminRequest(name string, address string) *MsgRegisterAdminRequest {

	return &MsgRegisterAdminRequest{
		Address: address,
		Name:    name,
	}
}

func (msg MsgRegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	if msg.Address == "" {
		return ErrAdminDetailsNil
	}
	if msg.Name == "" {
		return ErrAdminNameNil
	}

	return nil
}

func (msg MsgRegisterAdminRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgRegisterAdminRequest) GetSigners() []sdk.AccAddress {
	sign := msg.SignerAddress
	addr, _ := sdk.AccAddressFromBech32(sign)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________

func NewMsgAcceptLeaveRequest(admin string, student string, status LeaveStatus, signer string) *MsgAcceptLeaveRequest {

	return &MsgAcceptLeaveRequest{
		Admin:         admin,
		Student:       student,
		Status:        status,
		SignerAddress: signer,
	}
}

func (msg MsgAcceptLeaveRequest) ValidateBasic() error {
	eg := msg.Admin
	if _, err := sdk.AccAddressFromBech32(eg); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}

	if msg.Student == "" {
		return ErrStudentDetailsNil
	}
	return nil
}

func (msg MsgAcceptLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgAcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	sign := msg.SignerAddress
	addr, _ := sdk.AccAddressFromBech32(sign)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________
