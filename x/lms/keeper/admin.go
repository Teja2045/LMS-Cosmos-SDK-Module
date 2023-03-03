package keeper

import (
	"fmt"
	"lmsmodule/x/lms/types"
	"log"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AdminRegister(ctx sdk.Context, req *types.MsgRegisterAdminRequest) error {

	store := ctx.KVStore(k.storeKey)
	val, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	}

	store.Set(types.AdminstoreKey(req.Address), val)
	return nil
}

func (k Keeper) GetAdmin(ctx sdk.Context, req string) []byte {

	// if the address is not cosmos address
	if _, err := sdk.AccAddressFromBech32(req); err != nil {
		log.Fatal(err)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.AdminstoreKey(req))
}

//----------------------------------------------------------------------

func (k Keeper) Accept(ctx sdk.Context, studentAddress string, adminAddress string) error {

	store := ctx.KVStore(k.storeKey)

	// if admin is not a registered admin
	if store.Get(types.AdminstoreKey(adminAddress)) == nil {
		return types.ErrAdminDoesNotExist
	}

	// if student is not registered
	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return types.ErrStudentDoesNotExist
	}

	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(studentAddress))))

	// if a student never applied for a leave
	if leaveId == 0 {
		return types.ErrLeaveNeverApplied
	}

	// if most recent leave request is already handled
	if store.Get(types.LeaveStatusStoreKey(studentAddress, leaveId)) != nil {
		return types.ErrLeaveAlreadyHandled
	}

	leaveStatus := &types.MsgAcceptLeaveRequest{
		Admin:   adminAddress,
		Student: studentAddress,
		Status:  types.LeaveStatus_STATUS_ACCEPTED,
	}

	val, _ := k.cdc.Marshal(leaveStatus)
	store.Set(types.LeaveStatusStoreKey(studentAddress, leaveId), val)

	var leave types.MsgAcceptLeaveRequest

	if err := k.cdc.Unmarshal(store.Get(types.LeaveStatusStoreKey(studentAddress, leaveId)), &leave); err != nil {
		panic(err)
	}

	// after handling the leave, the student is removed from the pending list
	pendingLeaveStudentList := DecodeList(store.Get(types.PendingLeaveStudentsStoreKey()))
	for i, student := range pendingLeaveStudentList {
		if student == studentAddress {
			pendingLeaveStudentList = append(pendingLeaveStudentList[:i], pendingLeaveStudentList[i+1:]...)
			store.Set(types.PendingLeaveStudentsStoreKey(), EncodeList(pendingLeaveStudentList))
			break
		}
	}
	return nil
}

//----------------------------------------------------------------------------

// reject is also very similar to accept method
func (k Keeper) Reject(ctx sdk.Context, studentAddress string, adminAddress string) error {
	if _, err := sdk.AccAddressFromBech32(adminAddress); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	if store.Get(types.AdminstoreKey(adminAddress)) == nil {
		return types.ErrAdminDoesNotExist
	}

	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return types.ErrStudentDoesNotExist
	}

	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(studentAddress))))
	if leaveId == 0 {
		return types.ErrLeaveNeverApplied
	}

	if store.Get(types.LeaveStatusStoreKey(studentAddress, leaveId)) == nil {
		return types.ErrLeaveAlreadyHandled
	}

	leaveStatus := &types.MsgAcceptLeaveRequest{
		Admin:   adminAddress,
		Student: studentAddress,
		Status:  types.LeaveStatus_STATUS_REJECTED,
	}

	val, _ := k.cdc.Marshal(leaveStatus)

	store.Set(types.LeaveStatusStoreKey(adminAddress, leaveId), val)
	pendingLeaveStudentList := DecodeList(store.Get(types.PendingLeaveStudentsStoreKey()))
	for i, student := range pendingLeaveStudentList {
		if student == studentAddress {
			pendingLeaveStudentList = append(pendingLeaveStudentList[:i], pendingLeaveStudentList[i+1:]...)
			store.Set(types.PendingLeaveStudentsStoreKey(), EncodeList(pendingLeaveStudentList))
			break
		}
	}
	return nil
}

func (k Keeper) ListPendingLeaveStudents(ctx sdk.Context, adminAddress string) ([]string, error) {
	if _, err := sdk.AccAddressFromBech32(adminAddress); err != nil {
		fmt.Println("___here in admin register, error___")
		return []string{}, err
	}

	if adminAddress == "" {
		return []string{}, types.ErrAdminNameNil
	}

	store := ctx.KVStore(k.storeKey)

	if store.Get(types.AdminstoreKey(adminAddress)) == nil {
		return []string{}, types.ErrAdminDoesNotExist
	}
	StudentAddressList := store.Get(types.PendingLeaveStudentsStoreKey())
	return DecodeList(StudentAddressList), nil
}
