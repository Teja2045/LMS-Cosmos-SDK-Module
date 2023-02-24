package keeper

import (
	"fmt"
	"lmsmodule/x/lms/types"
	"log"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddStudent(ctx sdk.Context, student *types.Student) error {
	if _, err := sdk.AccAddressFromBech32(student.Address); err != nil {
		fmt.Println("___here in admin register, error___")
		return err
	}

	if student.Name == "" {
		return types.ErrStudentNameNil
	}

	if student.Id == "" {
		return types.ErrStudentIdNil
	}

	store := ctx.KVStore(k.storeKey)
	val, err := k.cdc.Marshal(student)
	if err != nil {
		return err
	} else {
		if store.Get(types.StudentStoreKey(student.Address)) != nil {
			return types.ErrStudentAlreadyExists
		}
		store.Set(types.StudentStoreKey(student.Address), val)
	}

	return nil
}

func (k Keeper) GetStudent(ctx sdk.Context, studentaddress string) []byte {
	if _, err := sdk.AccAddressFromBech32(studentaddress); err != nil {
		log.Fatal(err)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.StudentStoreKey(studentaddress))
}

func (k Keeper) CheckLeaveStatus(ctx sdk.Context, studentAddress string) (types.Leave, error) {
	if _, err := sdk.AccAddressFromBech32(studentAddress); err != nil {
		return types.Leave{}, err
	}

	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return types.Leave{}, types.ErrStudentDoesNotExist
	}

	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(studentAddress))))

	// if a student never applied for a leave
	if leaveId == 0 {
		return types.Leave{}, types.ErrLeaveNeverApplied
	}

	val := store.Get(types.LeaveStatusStoreKey(studentAddress, leaveId))
	var leave *types.MsgAcceptLeaveRequest
	k.cdc.Unmarshal(val, leave)
	return types.Leave{
		Address: leave.Admin,
		Status:  (leave.Status == types.LeaveStatus_STATUS_ACCEPTED),
	}, nil
}
