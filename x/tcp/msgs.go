package tcp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


// MsgTransfer defines a transfer message
type MsgTransfer struct {
	From sdk.AccAddress
	To sdk.AccAddress
	Value sdk.Coin
	// State []byte // TODO
	// Fee sdk.Coin

	// validation
	// From balance >= Value + Fee
	// From balance - Value - Fee = NewState of From
	// To balance + Value = NewState of To
}


// MsgContractDeploy defines a ContractDeploy message
type MsgContractDeploy struct {
	From sdk.AccAddress
	CID sdk.AccAddress
	Code []byte
	CodeHash string
	State []byte // TODO
	Fee sdk.Coin
}

// MsgContractExec defines a ontractExec message
type MsgContractExec struct {
	From sdk.AccAddress
	CID sdk.AccAddress
	State []byte // TODO
	RequestParams []RequestParam // TODO
	Proof []byte // TODO
	ResultHash string
	Fee sdk.Coin
}



// NewMsgTransfer is a constructor function for MsgTransfer
func NewMsgTransfer(from sdk.AccAddress, to sdk.AccAddress, value sdk.Coin, fee sdk.Coin) MsgTransfer {
	return MsgTransfer{
		From :  from,
		To : to,
		Value : value,
		Fee : fee
	}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgTransfer) ValidateBasic() sdk.Error {
	if msg.From.Empty()) {
		return sdk.ErrInvalidAddress(msg.From.String())
	}

	if msg.To.Empty()) {
		return sdk.ErrInvalidAddress(msg.To.String())
	}

	if msg.Value.IsNegative() {
		return sdk.ErrUnknownRequest("Transfer Value cannot be negative")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransfer) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}


// NewMsgContractDeploy is a constructor function for MsgTransfer
func NewMsgContractDeploy(from sdk.AccAddress) MsgContractDeploy {
	// TODO
	return MsgContractDeploy{
		From :  from
	}
}

// NewMsgContractDeploy is a constructor function for MsgTransfer
func MsgContractExec(from sdk.AccAddress) MsgContractExec {
	// TODO
	return MsgContractExec{
		From :  from
	}
}
