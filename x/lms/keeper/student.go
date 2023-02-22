package keeper

import (
	"fmt"
	"lmsmodule/x/lms/types"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddStudent(ctx sdk.Context, student *types.Student) error {
	if _, err := sdk.AccAddressFromBech32(student.Address); err != nil {
		fmt.Println("___here in admin register, error___")
		return err
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

	//fmt.Println("inside keeper.go after storing ", store.Get(types.AdminstoreKey(req.Address)))
	return nil
}

func (k Keeper) GetStudent(ctx sdk.Context, req string) []byte {
	if _, err := sdk.AccAddressFromBech32(req); err != nil {
		log.Fatal(err)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.StudentStoreKey(req))
}
