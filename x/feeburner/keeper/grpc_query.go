package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/oxygene76/medasdigital/x/feeburner/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TotalBurnedMedasAmount(goCtx context.Context, _ *types.QueryTotalBurnedMedasAmountRequest) (*types.QueryTotalBurnedMedasAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	totalBurnedMedasAmount := k.GetTotalBurnedMedasAmount(ctx)

	return &types.QueryTotalBurnedMedasAmountResponse{TotalBurnedMedasAmount: totalBurnedMedasAmount}, nil
}
