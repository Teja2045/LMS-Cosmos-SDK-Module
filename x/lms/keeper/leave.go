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

func (k Keeper) GetLeave(ctx sdk.Context, studentAddress string) (*types.MsgApplyLeaveRequest, error) {
	//fmt.Println("in GetLeave1: ")
	if _, err := sdk.AccAddressFromBech32(studentAddress); err != nil {
		return &types.MsgApplyLeaveRequest{}, err
	}
	//fmt.Println("in GetLeave2: ")
	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return &types.MsgApplyLeaveRequest{}, types.ErrStudentDoesNotExist
	}
	//fmt.Println("in GetLeave3: ")
	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(studentAddress))))
	if leaveId == 0 {
		return &types.MsgApplyLeaveRequest{}, types.ErrLeaveNeverApplied
	}
	//fmt.Println("in GetLeave4: ")
	leave := &types.MsgApplyLeaveRequest{}
	k.cdc.Unmarshal(store.Get(types.LeaveStoreKey(studentAddress, leaveId)), leave)
	fmt.Println(studentAddress, " ", leave)
	return leave, nil
}
