package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"

	"medasdigital/app/params"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMedasDenom        = []byte("MedasDenom")
	DefaultMedasDenom    = params.DefaultDenom
	KeyReserveAddress      = []byte("ReserveAddress")
	DefaultReserveAddress  = ""
	KeyTreasuryAddress     = []byte("TreasuryAddress")
	DefaultTreasuryAddress = ""
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(
			KeyMedasDenom,
			DefaultMedasDenom,
			validateMedasDenom,
		),
		paramtypes.NewParamSetPair(
			KeyReserveAddress,
			DefaultReserveAddress,
			func(i interface{}) error { return nil },
		),
		paramtypes.NewParamSetPair(
			KeyTreasuryAddress,
			DefaultTreasuryAddress,
			validateTreasuryAddress,
		),
	)
}

// NewParams creates a new Params instance
func NewParams(medasDenom, treasuryAddress string) Params {
	return Params{
		MedasDenom:    medasDenom,
		ReserveAddress:  DefaultReserveAddress,
		TreasuryAddress: treasuryAddress,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMedasDenom, DefaultTreasuryAddress)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyMedasDenom,
			&p.MedasDenom,
			validateMedasDenom,
		),
		paramtypes.NewParamSetPair(
			KeyReserveAddress,
			&p.ReserveAddress,
			func(i interface{}) error { return nil },
		),
		paramtypes.NewParamSetPair(
			KeyTreasuryAddress,
			&p.TreasuryAddress,
			validateTreasuryAddress,
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateMedasDenom(p.MedasDenom)
	if err != nil {
		return err
	}

	err = validateTreasuryAddress(p.TreasuryAddress)
	if err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMedasDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return fmt.Errorf("MedasDenom must not be empty")
	}

	return nil
}

func validateTreasuryAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Reserve might be explicitly empty in test environments
	if len(v) == 0 {
		return nil
	}

	_, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return fmt.Errorf("invalid Reserve address: %w", err)
	}

	return nil
}
