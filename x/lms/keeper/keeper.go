package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Keeper of the lms store
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

func (k Keeper) StoreKey() storetypes.StoreKey {
	return k.storeKey
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

func (k Keeper) GetCodec() codec.BinaryCodec {
	return k.cdc
}

func (k Keeper) GetStoreKey() storetypes.StoreKey {
	return k.storeKey
}
