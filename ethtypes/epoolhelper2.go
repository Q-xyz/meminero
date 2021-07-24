// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethtypes

import (
	"math/big"

	web3types "github.com/alethio/web3-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lacasian/ethwheels/ethgen"
)

// Reference imports to suppress errors
var (
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = web3types.Log{}
)

const Epoolhelper2ABI = "[{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"}],\"name\":\"currentRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"}],\"name\":\"delta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deltaA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deltaB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rDiv\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"eTokenForTokenATokenB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"feeAFeeBForEToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"tokenAForTokenB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tokenATokenBForEToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalA\",\"type\":\"uint256\"}],\"name\":\"tokenATokenBForTokenA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalB\",\"type\":\"uint256\"}],\"name\":\"tokenATokenBForTokenB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"name\":\"tokenBForTokenA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"totalA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"totalB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var Epoolhelper2 = NewEpoolhelper2Decoder()

type Epoolhelper2Decoder struct {
	*ethgen.Decoder
}

func NewEpoolhelper2Decoder() *Epoolhelper2Decoder {
	dec := ethgen.NewDecoder(Epoolhelper2ABI)
	return &Epoolhelper2Decoder{
		dec,
	}
}
