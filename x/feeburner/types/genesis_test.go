package types_test

import (
	"testing"

	"github.com/oxygene76/medasdigital/app"

	"github.com/stretchr/testify/require"

	keepertest "github.com/oxygene76/medasdigital/testutil/feeburner/keeper"
	"github.com/oxygene76/medasdigital/testutil/feeburner/nullify"
	"github.com/oxygene76/medasdigital/x/feeburner"
	"github.com/oxygene76/medasdigital/x/feeburner/types"
)

func TestGenesis(t *testing.T) {
	_ = app.GetDefaultConfig()

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

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "empty medas denom",
			genState: &types.GenesisState{
				Params: types.Params{
					MedasDenom: "",
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
