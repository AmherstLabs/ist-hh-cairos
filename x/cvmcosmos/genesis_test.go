package cvmcosmos_test

import (
	"testing"

	keepertest "cvm-cosmos/testutil/keeper"
	"cvm-cosmos/testutil/nullify"
	"cvm-cosmos/x/cvmcosmos"
	"cvm-cosmos/x/cvmcosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CvmcosmosKeeper(t)
	cvmcosmos.InitGenesis(ctx, *k, genesisState)
	got := cvmcosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
