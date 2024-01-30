package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	OpBNBUrl             string
	OrderContractAddress string
	OpBNBTokenUsdt       string
	SK                   string
	Wallet               string
	Leverage             uint
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
	leverageStr := os.Getenv("Leverage")
	leverageInt, err := strconv.Atoi(leverageStr)
	if err != nil {
		log.Fatal(err)
	}

	if leverageInt < 5 {
		log.Fatal("Minimum leverage of 5x is required")
	}

	if leverageInt >= 100 {
		log.Fatal("leverage is too high")
	}

	Leverage = uint(leverageInt)
}
