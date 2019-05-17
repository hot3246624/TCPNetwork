package tcp

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type ConAccount struct {

	Account auth.BaseAccount	`json:"account"`
	Code []byte					`json:"code"`
	CodeHash []byte				`json:"code_hash"`
	Result map[string]string	`json:"result"`
}

func NewTCPWithDeploy(caller sdk.AccAddress,contractCode []byte) ConAccount{

	addr := caller
	nonce := uint64(8)
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