package feeburner_test

import (
	"testing"

	keepertest "medasdigital/testutil/keeper"
	"medasdigital/testutil/nullify"

	"medasdigital/x/feeburner"
	"medasdigital/x/feeburner/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.FeeburnerKeeper(t)
	feeburner.InitGenesis(ctx, *k, genesisState)
	got := feeburner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
