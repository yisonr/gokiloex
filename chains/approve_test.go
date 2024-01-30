package chains

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestApprove(t *testing.T) {
	usdtToken := common.HexToAddress(OpBNBTokenUsdt)

	approveAddress := common.HexToAddress(OrderContractAddress)
	amount := big.NewInt(5000000000) // usdt decimal=8

	var approveUsdt ApproveSt
	approveUsdt.TokenAddress = &usdtToken
	approveUsdt.ApproveAddress = &approveAddress
	approveUsdt.Amount = amount

	tx, err := Approve(approveUsdt)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("approve success: %s\n", tx)

	ok, err := CheckTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Log("transaction confirmed")
	} else {
		t.Log("transaction failed")
	}
}
