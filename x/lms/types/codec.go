package types

import (
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// cdc.RegisterConcrete(&MsgRegisterAdminRequest{}, "lmsmodule/MsgRegisterAdminRequest", nil)
	// cdc.RegisterConcrete(&MsgAddStudentRequest{}, "lmsmodule/MsgAddStudentRequest", nil)
	// cdc.RegisterConcrete(&MsgApplyLeaveRequest{}, "lmsmodule/MsgApplyLeaveRequest", nil)
	// cdc.RegisterConcrete(&MsgAcceptLeaveRequest{}, "lmsmodule/MsgAcceptLeaveRequest", nil)
	legacy.RegisterAminoMsg(cdc, &MsgRegisterAdminRequest{}, "lmsmodule/MsgRegisterAdminRequest")
	legacy.RegisterAminoMsg(cdc, &MsgAddStudentRequest{}, "lmsmodule/MsgAddStudentRequest")
	legacy.RegisterAminoMsg(cdc, &MsgApplyLeaveRequest{}, "lmsmodule/MsgApplyLeaveRequest")
	legacy.RegisterAminoMsg(cdc, &MsgAcceptLeaveRequest{}, "lmsmodule/MsgAcceptLeaveRequest")
}
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterAdminRequest{},
		&MsgAddStudentRequest{},
		&MsgApplyLeaveRequest{},
		&MsgAcceptLeaveRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
