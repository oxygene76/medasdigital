package keeper_test

import (
	"testing"

	testkeeper "github.com/oxygene/medasdigital/testutil/keeper"
	"github.com/oxygene/medasdigital/x/medasdigital/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MedasdigitalKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
