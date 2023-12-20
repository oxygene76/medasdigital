package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	feekeeperutil "medasdigital/testutil/feeburner/keeper"
	"medasdigital/x/feeburner/types"
)

func TestGrpcQuery_TotalBurnedMedasAmount(t *testing.T) {
	feeKeeper, ctx := feekeeperutil.FeeburnerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	feeKeeper.RecordBurnedFees(ctx, sdk.NewCoin(types.DefaultMedasDenom, sdk.NewInt(100)))

	request := types.QueryTotalBurnedMedasAmountRequest{}
	response, err := feeKeeper.TotalBurnedMedasAmount(wctx, &request)
	require.NoError(t, err)
	require.Equal(t, &types.QueryTotalBurnedMedasAmountResponse{TotalBurnedMedasAmount: types.TotalBurnedMedasAmount{Coin: sdk.Coin{Denom: types.DefaultMedasDenom, Amount: sdk.NewInt(100)}}}, response)
}
