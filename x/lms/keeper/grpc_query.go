package keeper

import (
	"context"
	"lmsmodule/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListLeaves(ctx context.Context, listLeavesRequest *types.ListLeavesRequest) (*types.ListLeavesResponse, error) {
	return &types.ListLeavesResponse{}, nil
}
func (k Keeper) LeaveStatus(ctx context.Context, leaveStatusRequest *types.LeaveStatusRequest) (*types.LeaveStatusResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leave, _ := k.CheckLeaveStatus(sdkctx, leaveStatusRequest.Address)
	return &types.LeaveStatusResponse{Leave: &leave}, nil
}
