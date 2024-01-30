package config

import "math/big"

var PlaceOrderMethodIdWithTpAndSl = "655b3889" // Tp: take profit Sl: stop loss
var PlaceOrderMethodId = "00a30f93"            // 挂单
var MarketOrderMethodId = "fc2ee01d"           // 撤单
var ApproveMethodId = "095ea7b3"
var BaseBalance = big.NewInt(1000000000) // 10u

var (
	MaxPriorityFeePerGas        = big.NewInt(10000)
	MaxFeePerGas                = big.NewInt(10024)
	GasLimit             uint64 = 50000
)
