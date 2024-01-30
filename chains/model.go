package chains

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ApproveSt struct {
	TokenAddress   *common.Address
	ApproveAddress *common.Address
	Amount         *big.Int
}

type TransactionSt struct {
	SK                   string
	ContractAddress      *common.Address
	MethodId             [4]byte
	MaxPriorityFeePerGas *big.Int
	MaxFeePerGas         *big.Int
	GasLimit             uint64
	Value                *big.Int
	Data                 []byte
}
