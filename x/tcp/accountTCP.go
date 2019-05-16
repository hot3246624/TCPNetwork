package tcp

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"encoding/binary"
)

type ConAccount struct {
	Account auth.BaseAccount `json:"account"`
	Code []byte `json:"code"`
	CodeHash []byte `json:"codehash"`
	Result map[string]string `json:"result"`
}

func NewTCPWithDeploy(caller sdk.AccAddress,contractCode []byte) ConAccount{
	//for test, create a new account
	callerAcc := auth.NewBaseAccountWithAddress(caller)
	addr := callerAcc.GetAddress().Bytes()
	nonce := callerAcc.GetSequence()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, nonce)
	cAddr := append(addr, b...)
	contractAddr := sdk.AccAddress(cAddr)

	account := auth.NewBaseAccountWithAddress(contractAddr)
	account.SetSequence(0)
	ContractAcc := ConAccount{
		Account:account,
		Code:contractCode,
		CodeHash: contractAddr,
	}

	return ContractAcc
}