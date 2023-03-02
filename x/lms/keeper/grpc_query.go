package keeper

import (
	"context"
	"fmt"
	"lmsmodule/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListLeaves(ctx context.Context, listLeavesRequest *types.ListLeavesRequest) (*types.ListLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	studentAddressList, error := k.ListPendingLeaveStudents(sdkctx, listLeavesRequest.Address)
	leaves := []*types.Leave{}
	for _, studentAddress := range studentAddressList {
		leave1, _ := k.GetLeave(sdkctx, studentAddress)
		leaves = append(leaves, &types.Leave{Address: leave1.Address, Reason: leave1.Reason, Status: false})
	}
	return &types.ListLeavesResponse{Leaves: leaves}, error
}

func (k Keeper) LeaveStatus(ctx context.Context, leaveStatusRequest *types.LeaveStatusRequest) (*types.LeaveStatusResponse, error) {

	// fmt.Println("leaves in query")
	sdkctx := sdk.UnwrapSDKContext(ctx)
	// panic(fmt.Sprintf("value =================================%v", leaveStatusRequest))
	leave, err := k.CheckLeaveStatus(sdkctx, leaveStatusRequest.Address)
	// panic(fmt.Sprintf("at line1", leave, err))
	fmt.Println(leave, err)
	return &types.LeaveStatusResponse{Leave: &leave}, nil
	// return &types.LeaveStatusResponse{}, nil
}
