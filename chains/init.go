package chains

import (
	"gokiloex/config"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func init() {
	client, err := ethclient.Dial(config.OpBNBUrl)
	if err != nil {
		log.Fatal(err)
	}

	EthClient = client
}
