package keeper

import (
	"lmsmodule/x/lms/types"
	"log"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddStudent(ctx sdk.Context, student *types.Student) error {

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

//----------------------------------------------------------------------------

func (k Keeper) GetStudent(ctx sdk.Context, studentaddress string) *types.Student {
	if _, err := sdk.AccAddressFromBech32(studentaddress); err != nil {
		log.Fatal(err)
	}
	store := ctx.KVStore(k.storeKey)
	student := &types.Student{}
	k.cdc.Unmarshal(store.Get(types.StudentStoreKey(studentaddress)), student)

	return student
}

//----------------------------------------------------------------------------

func (k Keeper) CheckLeaveStatus(ctx sdk.Context, studentAddress string) (types.LeaveStatusResponse, error) {
	if _, err := sdk.AccAddressFromBech32(studentAddress); err != nil {
		return types.LeaveStatusResponse{}, err
	}

	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return types.LeaveStatusResponse{}, types.ErrStudentDoesNotExist
	}

	leaveIdInBytes := store.Get(types.LeaveCounterStoreKey(studentAddress))

	leaveId := 0

	if leaveIdInBytes != nil {
		leaveId, _ = strconv.Atoi(string(leaveIdInBytes))
	}

	// if a student never applied for a leave
	if leaveId == 0 {
		return types.LeaveStatusResponse{}, types.ErrLeaveNeverApplied
	}

	marshalledLeave := store.Get(types.LeaveStatusStoreKey(studentAddress, leaveId))
	if marshalledLeave == nil {
		return types.LeaveStatusResponse{
			SignedBy: "no admin handled it yet",
			Status:   types.LeaveStatus_STATUS_UNDEFINED,
		}, nil
	}
	var handledLeave types.MsgAcceptLeaveRequest
	k.cdc.Unmarshal(marshalledLeave, &handledLeave)

	leaveStatus := types.LeaveStatusResponse{
		SignedBy: handledLeave.Admin,
		Status:   handledLeave.Status,
	}
	return leaveStatus, nil
}

//----------------------------------------------------------------------------

func (k Keeper) GetStudents(ctx sdk.Context) []*types.Student {
	store := ctx.KVStore(k.storeKey)

	var students []*types.Student
	itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	for ; itr.Valid(); itr.Next() {
		var t types.Student
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)
	}
	return students
}
