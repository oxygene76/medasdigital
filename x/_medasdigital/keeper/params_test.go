package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "medasdigital/testutil/keeper"
	"medasdigital/x/medasdigital/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MedasdigitalKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
