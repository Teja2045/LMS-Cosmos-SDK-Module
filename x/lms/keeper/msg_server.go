package keeper

// "clms/x/lms/keeper"

// "github.com/cosmos/cosmos-sdk/x/staking/types"
// v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
// "github.com/cosmos/cosmos-sdk/x/bank/types"
// "github.com/cosmos/cosmos-sdk/x/bank/types"

import (
	"context"
	"lmsmodule/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type msgServer struct {
// 	Keeper
// }

var _ types.MsgServer = Keeper{}

// func NewMsgServerImpl(keeper Keeper) types.MsgServer {
// 	return &msgServer{Keeper: keeper}
// }

// func (k msgServer) mustEmbedUnimplementedMsgServer() {} when just used grpc protoc

func (k Keeper) MsgRegisterAdmin(ctx context.Context, req *types.MsgRegisterAdminRequest) (*types.MsgRegisterAdminResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	return &types.MsgRegisterAdminResponse{}, k.AdminRegister(sdkctx, req)
}

func (k Keeper) MsgAddStudent(ctx context.Context, req *types.MsgAddStudentRequest) (*types.MsgAddStudentResponse, error) {

	sdkctx := sdk.UnwrapSDKContext(ctx)
	for _, student := range req.Students {
		k.AddStudent(sdkctx, student)
	}
	return &types.MsgAddStudentResponse{}, nil
}

func (k Keeper) MsgApplyLeave(ctx context.Context, req *types.MsgApplyLeaveRequest) (*types.MsgApplyLeaveResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	return &types.MsgApplyLeaveResponse{}, k.AddLeave(sdkctx, req)
}

func (k Keeper) MsgAcceptLeave(ctx context.Context, req *types.MsgAcceptLeaveRequest) (*types.MsgAcceptLeaveResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	return &types.MsgAcceptLeaveResponse{}, k.Accept(sdkctx, req.Student, req.Admin)
}

// type msgServer struct {
// 	Keeper
// }

// // var _ types.MsgServer = msgServer{}

// var _ types.MsgServer = msgServer{}

// func (k msgServer) AddStudent(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
// 	return &types.AddStudentResponse{}, nil
// }

// func (k msgServer) RegisterAdmin(goctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goctx)
// 	Keeper.AdminRegister(ctx, req)
// 	return &types.RegisterAdminResponse{}, nil
// }
// func (k msgServer) ApplyLeave(goctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
// 	// k.AcceptLeave(ctx, req)
// 	ctx := sdk.UnwrapSDKContext(goctx)
// 	Keeper.AcceptLeaves(ctx, req)
// 	return &types.ApplyLeaveResponse{}, nil
// }
// func (k msgServer) AcceptLeave(goctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goctx)
// 	return &types.AcceptLeaveResponse{}, nil
// }
