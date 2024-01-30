package chains

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client
var (
	OpBNBUrl             string
	OrderContractAddress string
	OpBNBTokenUsdt       string
	SK                   string
	Wallet               string
)

func init() {
	OpBNBUrl = os.Getenv("OpBNBUrl")
	OrderContractAddress = os.Getenv("OrderContractAddress")
	OpBNBTokenUsdt = os.Getenv("OpBNBTokenUsdt")
	SK = os.Getenv("SK")
	Wallet = os.Getenv("Wallet")

	client, err := ethclient.Dial(OpBNBUrl)
	if err != nil {
		log.Fatal(err)
	}

	EthClient = client
}
