package keeper

import (
	"lmsmodule/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {

}

// ExportGenesis returns a GenesisState for a given context.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {

	return &types.GenesisState{}
}
