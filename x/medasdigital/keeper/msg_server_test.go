package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/oxygene/medasdigital/testutil/keeper"
	"github.com/oxygene/medasdigital/x/medasdigital/keeper"
	"github.com/oxygene/medasdigital/x/medasdigital/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MedasdigitalKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
