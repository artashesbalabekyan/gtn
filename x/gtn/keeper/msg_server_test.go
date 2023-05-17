package keeper_test

import (
	"context"
	"testing"

	keepertest "GTN/testutil/keeper"
	"GTN/x/gtn/keeper"
	"GTN/x/gtn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GtnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
