package keeper

import (
	"cvm-cosmos/x/cvmcosmos/types"
)

var _ types.QueryServer = Keeper{}
