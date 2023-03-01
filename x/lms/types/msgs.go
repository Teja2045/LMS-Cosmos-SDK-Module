package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgAddStudentRequest{}
	_ sdk.Msg = &MsgRegisterAdminRequest{}
	_ sdk.Msg = &MsgApplyLeaveRequest{}
	_ sdk.Msg = &MsgAcceptLeaveRequest{}
)

func NewMsgAddStudentRequest() *MsgAddStudentRequest {

	return &MsgAddStudentRequest{}
}

func (msg MsgAddStudentRequest) ValidateBasic() error {
	eg := msg.Admin
	if _, err := sdk.AccAddressFromBech32(eg); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	return nil
}

func (msg MsgAddStudentRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgAddStudentRequest) GetSigners() []sdk.AccAddress {
	sign := msg.Admin
	addr, _ := sdk.AccAddressFromBech32(sign)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________

func NewMsgApplyLeaveRequest() *MsgApplyLeaveRequest {

	return &MsgApplyLeaveRequest{}
}

func (msg MsgApplyLeaveRequest) ValidateBasic() error {
	eg := msg.Address
	if _, err := sdk.AccAddressFromBech32(eg); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	return nil
}

func (msg MsgApplyLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	sign := "some sign"
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
	eg := msg.Address
	if _, err := sdk.AccAddressFromBech32(eg); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	return nil
}

func (msg MsgRegisterAdminRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgRegisterAdminRequest) GetSigners() []sdk.AccAddress {
	sign := msg.Address
	addr, _ := sdk.AccAddressFromBech32(sign)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________

func NewMsgAcceptLeaveRequest() *MsgAcceptLeaveRequest {

	return &MsgAcceptLeaveRequest{}
}

func (msg MsgAcceptLeaveRequest) ValidateBasic() error {
	eg := msg.Admin
	if _, err := sdk.AccAddressFromBech32(eg); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
	}
	return nil
}

func (msg MsgAcceptLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgAcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	sign := msg.Admin
	addr, _ := sdk.AccAddressFromBech32(sign)
	return []sdk.AccAddress{addr}
}

//_________________________________________________________________________________________
