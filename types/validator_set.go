package types

import (
	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/go-crypto"

	"github.com/cosmos/cosmos-sdk/wire"
)

type Validator struct {
	Address Address       `json:"address"`
	PubKey  crypto.PubKey `json:"pub_key"`
	Power   Rat           `json:"voting_power"`
}

func (v Validator) ABCIValidator(cdc *wire.Codec) abci.Validator {
	pkBytes, err := cdc.MarshalBinary(v.PubKey)
	if err != nil {
		panic(err)
	}
	return abci.Validator{
		PubKey: pkBytes,
		Power:  v.Power.Evaluate(),
	}
}

func (v Validator) ABCIValidatorZero(cdc *wire.Codec) abci.Validator {
	pkBytes, err := cdc.MarshalBinary(v.PubKey)
	if err != nil {
		panic(err)
	}
	return abci.Validator{
		PubKey: pkBytes,
		Power:  0,
	}
}

type ValidatorSetKeeper interface {
	Hash(Context) []byte
	GetValidators(Context) []*Validator
	Size(Context) int
	IsValidator(Context, Address) bool
	GetByAddress(Context, Address) (int, *Validator)
	GetByIndex(Context, int) *Validator
	TotalPower(Context) Rat
}
