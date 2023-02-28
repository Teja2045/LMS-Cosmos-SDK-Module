package keeper

// "clms/x/lms/keeper"

// "github.com/cosmos/cosmos-sdk/x/staking/types"
// v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
// "github.com/cosmos/cosmos-sdk/x/bank/types"
// "github.com/cosmos/cosmos-sdk/x/bank/types"

import (
	"context"
	"lmsmodule/x/lms/types"
	"log"

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
	k.AdminRegister(sdkctx, req)
	return &types.MsgRegisterAdminResponse{}, nil
}

func (k Keeper) MsgAddStudent(ctx context.Context, req *types.MsgAddStudentRequest) (*types.MsgAddStudentResponse, error) {

	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		log.Fatal("___here in admin register, error___", err)
	}

	sdkctx := sdk.UnwrapSDKContext(ctx)
	for _, student := range req.Students {
		k.AddStudent(sdkctx, student)
	}
	return &types.MsgAddStudentResponse{}, nil
}

func (k Keeper) MsgApplyLeave(ctx context.Context, req *types.MsgApplyLeaveRequest) (*types.MsgApplyLeaveResponse, error) {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		log.Fatal("___here in admin register, error___", err)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	k.AddLeave(sdkctx, req)
	return &types.MsgApplyLeaveResponse{}, nil
}

func (k Keeper) MsgAcceptLeave(ctx context.Context, req *types.MsgAcceptLeaveRequest) (*types.MsgAcceptLeaveResponse, error) {
	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		log.Fatal("___here in admin register, error___", err)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	k.Accept(sdkctx, req.Student, req.Admin)
	return &types.MsgAcceptLeaveResponse{}, nil
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
