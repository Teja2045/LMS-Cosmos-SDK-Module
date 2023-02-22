package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgRegisterAdminRequest{}, "clms/MsgRegisterAdminRequest")
	legacy.RegisterAminoMsg(cdc, &MsgAddStudentRequest{}, "clms/MsgAddStudentRequest")
	legacy.RegisterAminoMsg(cdc, &MsgApplyLeaveRequest{}, "clms/MsgApplyLeaveRequest")
	legacy.RegisterAminoMsg(cdc, &MsgAcceptLeaveRequest{}, "clms/MsgApplyLeaveRequest")
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
	sdk.RegisterLegacyAminoCodec(amino)
}
