package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	OpBNBUrl             string
	OrderContractAddress string
	OpBNBTokenUsdt       string
	SK                   string
	Wallet               string
)

func init() {
	// 加载 .env 文件
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	OpBNBUrl = os.Getenv("OpBNBUrl")
	OrderContractAddress = os.Getenv("OrderContractAddress")
	OpBNBTokenUsdt = os.Getenv("OpBNBTokenUsdt")
	SK = os.Getenv("SK")
	Wallet = os.Getenv("Wallet")
}
