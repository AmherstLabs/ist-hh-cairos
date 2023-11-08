package keeper_test

import (
	"context"
	"testing"

	keepertest "cvm-cosmos/testutil/keeper"
	"cvm-cosmos/x/cvmcosmos/keeper"
	"cvm-cosmos/x/cvmcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CvmcosmosKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
