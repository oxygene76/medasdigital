package keeper

import (
	"fmt"

	"cosmossdk.io/errors"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	consumertypes "github.com/cosmos/interchain-security/v3/x/ccv/consumer/types"

	"medasdigital/x/feeburner/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		authority     string
	}
)

var KeyBurnedFees = []byte("BurnedFees")

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	authority string,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		authority:     authority,
	}
}

func (k Keeper) GetAuthority() string {
	return k.authority
}

// RecordBurnedFees adds `amount` to the total amount of burned MEDAS tokens
func (k Keeper) RecordBurnedFees(ctx sdk.Context, amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	totalBurnedMedasAmount := k.GetTotalBurnedMedasAmount(ctx)
	totalBurnedMedasAmount.Coin = totalBurnedMedasAmount.Coin.Add(amount)

	store.Set(KeyBurnedFees, k.cdc.MustMarshal(&totalBurnedMedasAmount))
}

// GetTotalBurnedMedasAmount gets the total burned amount of MEDAS tokens
func (k Keeper) GetTotalBurnedMedasAmount(ctx sdk.Context) types.TotalBurnedMedasAmount {
	store := ctx.KVStore(k.storeKey)

	var totalBurnedMedasAmount types.TotalBurnedMedasAmount
	bzTotalBurnedMedasAmount := store.Get(KeyBurnedFees)
	if bzTotalBurnedMedasAmount != nil {
		k.cdc.MustUnmarshal(bzTotalBurnedMedasAmount, &totalBurnedMedasAmount)
	}

	if totalBurnedMedasAmount.Coin.Denom == "" {
		totalBurnedMedasAmount.Coin = sdk.NewCoin(k.GetParams(ctx).MedasDenom, sdk.NewInt(0))
	}

	return totalBurnedMedasAmount
}

// BurnAndDistribute is an important part of tokenomics. It does few things:
// 1. Burns MEDAS fee coins distributed to consumertypes.ConsumerRedistributeName in ICS (https://github.com/cosmos/interchain-security/v3/blob/v0.2.0/x/ccv/consumer/keeper/distribution.go#L17)
// 2. Updates total amount of burned MEDAS coins
// 3. Sends non-MEDAS fee tokens to reserve contract address
// Panics if no `consumertypes.ConsumerRedistributeName` module found OR could not burn MEDAS tokens
func (k Keeper) BurnAndDistribute(ctx sdk.Context) {
	moduleAddr := k.accountKeeper.GetModuleAddress(consumertypes.ConsumerRedistributeName)
	if moduleAddr == nil {
		panic("ConsumerRedistributeName must have module address")
	}

	params := k.GetParams(ctx)
	balances := k.bankKeeper.GetAllBalances(ctx, moduleAddr)
	fundsForReserve := make(sdk.Coins, 0, len(balances))

	for _, balance := range balances {
		if !balance.IsZero() {
			if balance.Denom == params.MedasDenom {
				err := k.bankKeeper.BurnCoins(ctx, consumertypes.ConsumerRedistributeName, sdk.Coins{balance})
				if err != nil {
					panic(errors.Wrapf(err, "failed to burn MEDAS tokens during fee processing"))
				}

				k.RecordBurnedFees(ctx, balance)
			} else {
				fundsForReserve = append(fundsForReserve, balance)
			}
		}
	}

	if len(fundsForReserve) > 0 {
		addr, err := sdk.AccAddressFromBech32(params.TreasuryAddress)
		if err != nil {
			// there's no way we face this kind of situation in production, since it means the chain is misconfigured
			// still, in test environments it might be the case when the chain is started without Reserve
			// in such case we just burn the tokens
			err := k.bankKeeper.BurnCoins(ctx, consumertypes.ConsumerRedistributeName, fundsForReserve)
			if err != nil {
				panic(errors.Wrapf(err, "failed to burn tokens during fee processing"))
			}
		} else {
			err = k.bankKeeper.SendCoins(
				ctx,
				moduleAddr, addr,
				fundsForReserve,
			)
			if err != nil {
				panic(errors.Wrapf(err, "failed sending funds to Reserve"))
			}
		}
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// FundCommunityPool is method to satisfy DistributionKeeper interface for packet-forward-middleware Keeper.
// The original method sends coins to a community pool of a chain.
// The current method sends coins to a Fee Collector module which collects fee on consumer chain.
func (k Keeper) FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, authtypes.FeeCollectorName, amount)
}
