package chains

import (
	"gokiloex/config"
	"math/big"
)

func Approve(approve ApproveSt) (string, error) {
	// 构建 approve 交易data
	approveMethodBytes, err := ConvMethodIdToByte(config.ApproveMethodId)
	if err != nil {
		return "", err
	}
	contractBytes := approve.ApproveAddress.Bytes()
	amountBytes := approve.Amount.Bytes()

	data := BuildABIData(approveMethodBytes, contractBytes, amountBytes)

	var approveTrans = TransactionSt{
		SK:                   config.SK,
		ContractAddress:      approve.TokenAddress,
		MethodId:             approveMethodBytes,
		MaxPriorityFeePerGas: config.MaxPriorityFeePerGas,
		MaxFeePerGas:         config.MaxFeePerGas,
		GasLimit:             config.GasLimit,
		Value:                big.NewInt(0),
		Data:                 data,
	}

	tx, err := SendTransaction(approveTrans)
	return tx, err
}
