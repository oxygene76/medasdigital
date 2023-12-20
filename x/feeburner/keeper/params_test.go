package keeper_test

import (
	"testing"

	"medasdigital/app"

	"github.com/stretchr/testify/require"

	testkeeper "medasdigital/testutil/feeburner/keeper"
	"medasdigital/x/feeburner/types"
)

func TestGetParams(t *testing.T) {
	_ = app.GetDefaultConfig()

	k, ctx := testkeeper.FeeburnerKeeper(t)
	params := types.DefaultParams()

	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	require.EqualValues(t, params, k.GetParams(ctx))
}
