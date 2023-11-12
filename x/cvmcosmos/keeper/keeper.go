package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"cvm-cosmos/x/cvmcosmos/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetCairoContract(ctx sdk.Context, contractName string, data []byte) {
    store := ctx.KVStore(k.storeKey)
    key := []byte(contractName)  // Or use a hash function
    store.Set(key, data)
}

func (k Keeper) GetCairoContract(ctx sdk.Context, contractName string) ([]byte, bool) {
    store := ctx.KVStore(k.storeKey)
    key := []byte(contractName)  // Or use a hash function
    if !store.Has(key) {
        return nil, false
    }
    return store.Get(key), true
}