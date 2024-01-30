package chains

import (
	"math/big"
	"testing"
)

func TestGetBalance(t *testing.T) {
	balance, err := GetBalance(OpBNBTokenUsdt, Wallet)
	if err != nil {
		t.Fatal(err)
	}

	realBalance := big.NewInt(0)
	realBalance.SetString("55303093300000000000", 10)

	t.Log(realBalance.String())
	if balance.String() != realBalance.String() {
		t.Errorf("err balance: %s\n", balance.String())
	}
}
