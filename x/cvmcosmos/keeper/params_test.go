package keeper_test

import (
	"testing"

	testkeeper "cvm-cosmos/testutil/keeper"
	"cvm-cosmos/x/cvmcosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CvmcosmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
