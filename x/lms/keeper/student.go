package keeper

import (
	"lmsmodule/x/lms/types"
	"log"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddStudent(ctx sdk.Context, student *types.Student) error {
	// if _, err := sdk.AccAddressFromBech32(student.Address); err != nil {
	// 	log.Println(student.Address)
	// 	fmt.Println("___here in admin register, error___")
	// 	return types.ErrStudentAlreadyExists
	// }

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

func (k Keeper) GetStudent(ctx sdk.Context, studentaddress string) *types.Student {
	if _, err := sdk.AccAddressFromBech32(studentaddress); err != nil {
		log.Fatal(err)
	}
	store := ctx.KVStore(k.storeKey)
	student := &types.Student{}
	k.cdc.Unmarshal(store.Get(types.StudentStoreKey(studentaddress)), student)

	return student
}

func (k Keeper) CheckLeaveStatus(ctx sdk.Context, studentAddress string) (types.Leave, error) {
	if _, err := sdk.AccAddressFromBech32(studentAddress); err != nil {
		return types.Leave{}, err
	}

	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
		return types.Leave{}, types.ErrStudentDoesNotExist
	}

	leaveIdInBytes := store.Get(types.LeaveCounterStoreKey(studentAddress))

	leaveId := 0

	if leaveIdInBytes != nil {
		leaveId, _ = strconv.Atoi(string(leaveIdInBytes))
	}

	//panic("yoooooooooooooooooooo")

	// if a student never applied for a leave
	if leaveId == 0 {

		return types.Leave{}, types.ErrLeaveNeverApplied
	}

	//panic(fmt.Sprint(leaveId))

	//panic("yoooooooooooooooooooo")

	val := store.Get(types.LeaveStatusStoreKey(studentAddress, leaveId))
	if val == nil {
		return types.Leave{
			Address: "no admin handled it yet",
			Status:  false,
		}, nil
	}
	var leave types.MsgAcceptLeaveRequest
	//panic(fmt.Sprint("yoooooooooooooooooooo", val))
	k.cdc.Unmarshal(val, &leave)
	//panic(fmt.Sprint("sdjhd", err))

	res := types.Leave{
		Address: leave.Admin,
		Status:  true,
	}
	//panic(res)
	return res, nil
}
