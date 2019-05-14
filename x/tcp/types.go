package tcp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Amount struct {
	Address []sdk.AccAddress
	Value   []sdk.Coin
}

// user request for ContractExec
type RequestParam struct {
	From        sdk.AccAddress
	CID         sdk.AccAddress
	Proxy       sdk.AccAddress
	DataSources []Amount
	Fee         sdk.Coin
	Sig         []byte // signature

}

type Balance struct {
	Address  []sdk.AccAddress
	balances []sdk.Coin
}

type State struct {
	balances []Balance
}
