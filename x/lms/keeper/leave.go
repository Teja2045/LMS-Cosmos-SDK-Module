package keeper

import (
	"fmt"
	"lmsmodule/x/lms/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddLeave(ctx sdk.Context, leave *types.Leave) error {
	if _, err := sdk.AccAddressFromBech32(leave.Address); err != nil {
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
	leaveId, err := strconv.Atoi(string(leaveid))
	if err != nil {
		fmt.Println(err)
	}
	if leaveid == nil {
		leaveId = 0
	}

	leaveId++

	store.Set(types.LeaveCounterStoreKey(leave.Address), []byte(strconv.Itoa(leaveId)))

	store.Set(types.LeaveStoreKey(leave.Address, leaveId), val)

	pendingLeaveStudentList := DecodeList(store.Get(types.PendingLeaveStudentsStoreKey()))

	for _, studentAddress := range pendingLeaveStudentList {
		if studentAddress == leave.Address {
			return nil
		}
	}
	pendingLeaveStudentList = append(pendingLeaveStudentList, leave.Address)
	store.Set(types.PendingLeaveStudentsStoreKey(), EncodeList(pendingLeaveStudentList))
	return nil
}

//----------------------------------------------------------------------

func (k Keeper) GetLeave(ctx sdk.Context, studentAddress string) (*types.Leave, error) {
	if _, err := sdk.AccAddressFromBech32(studentAddress); err != nil {
		return &types.Leave{}, err
	}
	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return &types.Leave{}, types.ErrStudentDoesNotExist
	}
	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(studentAddress))))
	if leaveId == 0 {
		return &types.Leave{}, types.ErrLeaveNeverApplied
	}

	leave := types.Leave{}
	k.cdc.Unmarshal(store.Get(types.LeaveStoreKey(studentAddress, leaveId)), &leave)
	fmt.Println(studentAddress, " ", leave)
	return &leave, nil
}
