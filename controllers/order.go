package controllers

import (
	"fmt"
	"gokiloex/chains"
	"gokiloex/config"
)

func PlaceOrder() error {
	// 查询余额, TODO: 检查bnb
	balance, err := chains.GetBalance(config.OpBNBTokenUsdt, config.Wallet)
	if err != nil {
		return err
	}
	fmt.Println(balance)

	// if balance.Cmp(config.BaseBalance.Add(config.BaseBalance)) != 1 {
	// 	return errors.New("insufficent usdt balance")
	// }

	return nil
}
