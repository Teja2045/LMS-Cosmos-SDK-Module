package keeper

import (
	"context"
	"lmsmodule/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListPendingLeaves(ctx context.Context, listLeavesRequest *types.ListLeavesRequest) (*types.ListLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	studentAddressList, error := k.GetPendingLeaveStudents(sdkctx, listLeavesRequest.Address)
	leaves := []*types.Leave{}
	for _, studentAddress := range studentAddressList {
		leave1, _ := k.GetLeave(sdkctx, studentAddress)
		leaves = append(leaves, leave1)
	}
	return &types.ListLeavesResponse{Leaves: leaves}, error
}

func (k Keeper) LeaveStatus(ctx context.Context, leaveStatusRequest *types.LeaveStatusRequest) (*types.LeaveStatusResponse, error) {

	sdkctx := sdk.UnwrapSDKContext(ctx)
	leaveStatus, err := k.CheckLeaveStatus(sdkctx, leaveStatusRequest.Address)
	return &leaveStatus, err
}

func (k Keeper) ListStudents(ctx context.Context, listStudentsRequest *types.ListStudentsRequest) (*types.ListStudentsResponse, error) {

	sdkctx := sdk.UnwrapSDKContext(ctx)
	if k.CheckAdmin(sdkctx, listStudentsRequest.Address) {
		return &types.ListStudentsResponse{}, types.ErrAdminDoesNotExist
	}
	students := k.GetStudents(sdkctx)
	return &types.ListStudentsResponse{Students: students}, nil
}

func (k Keeper) ListHandledLeaves(ctx context.Context, listHandledLeavesRequest *types.ListHandledLeavesRequest) (*types.ListHandledLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leaves, error := k.GetHandledLeaves(sdkctx, listHandledLeavesRequest.Address)

	return &types.ListHandledLeavesResponse{Leaves: leaves}, error
}

func (k Keeper) ListAllAcceptedLeaves(ctx context.Context, listHandledLeavesRequest *types.ListHandledLeavesRequest) (*types.ListHandledLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leaves, error := k.GetAllAcceptedLeaves(sdkctx, listHandledLeavesRequest.Address)

	return &types.ListHandledLeavesResponse{Leaves: leaves}, error
}

func (k Keeper) ListAllRejectedLeaves(ctx context.Context, listHandledLeavesRequest *types.ListHandledLeavesRequest) (*types.ListHandledLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leaves, error := k.GetAllRejectedLeaves(sdkctx, listHandledLeavesRequest.Address)

	return &types.ListHandledLeavesResponse{Leaves: leaves}, error
}
