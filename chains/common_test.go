package chains

import (
	"kiloex/config"
	"math/big"
	"testing"
)

func TestGetBalance(t *testing.T) {
	balance, err := GetBalance(config.OpBNBTokenUsdt, config.Wallet)
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
