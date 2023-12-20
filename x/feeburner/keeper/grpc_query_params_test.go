package keeper_test

import (
	"testing"

	"github.com/oxygene76/medasdigital/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/oxygene76/medasdigital/testutil/feeburner/keeper"
	"github.com/oxygene76/medasdigital/x/feeburner/types"
)

func TestParamsQuery(t *testing.T) {
	_ = app.GetDefaultConfig()

	keeper, ctx := testkeeper.FeeburnerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	err := keeper.SetParams(ctx, params)
	require.NoError(t, err)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
