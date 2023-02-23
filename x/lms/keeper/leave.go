package keeper

import (
	"fmt"
	"lmsmodule/x/lms/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddLeave(ctx sdk.Context, leave *types.MsgApplyLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(leave.Address); err != nil {
		fmt.Println("___here in AddLeave, error___")
		return err
	}

	store := ctx.KVStore(k.storeKey)

	if store.Get(types.StudentStoreKey(leave.Address)) == nil {
		return types.ErrStudentDoesNotExist
	}
	val, err := k.cdc.Marshal(leave)
	if err != nil {
		return err
	}
	leaveid := store.Get(types.LeaveCounterStoreKey(leave.Address))
	leaveId, _ := strconv.Atoi(string(leaveid))
	//k.cdc.Unmarshal(leaveid, &ty)
	if leaveid == nil {
		leaveId = 0
	}
	leaveId++
	store.Set(types.LeaveCounterStoreKey(leave.Address), []byte(strconv.Itoa(leaveId)))
	store.Set(types.LeaveStoreKey(leave.Address, leaveId), val)

	return nil
}

func (k Keeper) GetLeave(ctx sdk.Context, req string) (*types.MsgApplyLeaveRequest, error) {
	if _, err := sdk.AccAddressFromBech32(req); err != nil {
		return &types.MsgApplyLeaveRequest{}, err
	}
	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(req)) == nil {
		return &types.MsgApplyLeaveRequest{}, types.ErrStudentDoesNotExist
	}
	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(req))))
	if leaveId == 0 {
		return &types.MsgApplyLeaveRequest{}, types.ErrLeaveNeverApplied
	}
	leave := &types.MsgApplyLeaveRequest{}
	k.cdc.Unmarshal(store.Get(types.LeaveStoreKey(req, leaveId)), leave)
	fmt.Println(req, " ", leave)
	return leave, nil
}
