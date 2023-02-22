package keeper

import (
	"context"
	"lmsmodule/x/lms/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListLeaves(context.Context, *types.ListLeavesRequest) (*types.ListLeavesResponse, error) {
	return &types.ListLeavesResponse{}, nil
}
func (k Keeper) LeaveStatus(context.Context, *types.LeaveStatusRequest) (*types.LeaveStatusResponse, error) {
	return &types.LeaveStatusResponse{}, nil
}
