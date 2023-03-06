package keeper

import (
	"context"
	"lmsmodule/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListLeaves(ctx context.Context, listLeavesRequest *types.ListLeavesRequest) (*types.ListLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	studentAddressList, error := k.ListPendingLeaveStudents(sdkctx, listLeavesRequest.Address)
	leaves := []*types.MsgApplyLeaveRequest{}
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
