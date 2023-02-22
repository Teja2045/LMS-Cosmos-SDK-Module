package keeper

import (
	"fmt"
	"lmsmodule/x/lms/types"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AdminRegister(ctx sdk.Context, req *types.MsgRegisterAdminRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		fmt.Println("___here in admin register, error___")
		return err
	}

	store := ctx.KVStore(k.storeKey)
	val, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.AdminstoreKey(req.Address), val)
	}
	//fmt.Println("inside keeper.go after storing ", store.Get(types.AdminstoreKey(req.Address)))
	return nil
}

func (k Keeper) GetAdmin(ctx sdk.Context, req string) []byte {
	if _, err := sdk.AccAddressFromBech32(req); err != nil {
		log.Fatal(err)
	}
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.AdminstoreKey(req))

}
