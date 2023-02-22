package keeper

import (
	"lmsmodule/x/lms/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the lms store
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new lms Keeper instance
func NewKeeper(key storetypes.StoreKey,
	cdc codec.BinaryCodec,
) Keeper {

	return Keeper{
		cdc:      cdc,
		storeKey: key,
	}
}

func (k Keeper) AdminRegister(ctx sdk.Context, req *types.MsgRegisterAdminRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.AdminKey, bz)
	}
	return nil
}

func (k Keeper) AcceptLeave(ctx sdk.Context, req *types.MsgAcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)

	req.Status = types.LeaveStatus_STATUS_ACCEPTED
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.LeaveKey, bz)
	}
	return nil
}
