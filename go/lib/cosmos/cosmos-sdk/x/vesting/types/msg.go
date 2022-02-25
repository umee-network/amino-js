package types

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgCreateVestingAccount{}
)

// MsgCreateVestingAccount defines a message that enables creating a vesting account.
// https://docs.cosmos.network/v0.42/core/proto-docs.html#cosmos.vesting.v1beta1.MsgCreateVestingAccount
type MsgCreateVestingAccount struct {
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
	Amount      sdk.Coins      `json:"amount"`
	EndTime     sdk.Int        `json:"end_time"`
	Delayed     bool           `json:"delayed"`
}
