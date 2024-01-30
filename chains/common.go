package chains

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func BuildABIData(methodId [4]byte, params ...([]byte)) []byte {
	var data []byte
	data = append(data, methodId[:]...)
	for _, param := range params {
		paramByte := common.LeftPadBytes(param, 32)
		data = append(data, paramByte...)
	}

	return data
}

func ConvMethodIdToByte(methodId string) ([4]byte, error) {
	selectorBytes, err := hex.DecodeString(methodId)
	if err != nil {
		// log.Fatal(err)
		return [4]byte{}, err
	}

	// 确保selectorBytes的长度是4
	if len(selectorBytes) != 4 {
		log.Fatalf("Invalid length of selector: got %d, want 4", len(selectorBytes))
		return [4]byte{}, errors.New("Invalid length of selector")
	}

	// 将字节切片转换为[4]byte数组
	var selector [4]byte
	copy(selector[:], selectorBytes[:4])

	// 打印转换后的结果
	// fmt.Printf("Selector: %x\n", selector)
	return selector, nil
}

func GetBalance(tokenAddress, walletAddress string) (*big.Int, error) {
	// token合约地址
	tokenContract := common.HexToAddress(tokenAddress)
	// 计算balanceOf函数的选择器
	funcSignature := []byte("balanceOf(address)")
	funcSelector := crypto.Keccak256Hash(funcSignature).Bytes()[:4]
	// fmt.Printf("Function Selector: %x\n", funcSelector)

	// 要查询的地址
	address := common.HexToAddress(walletAddress)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)

	// 构造调用数据
	data := append(funcSelector, paddedAddress...)

	// 发送请求
	balance, err := EthClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &tokenContract,
		Data: data,
	}, nil)
	if err != nil {
		// log.Fatalf("Failed to retrieve token balance: %v", err)
		return nil, err
	}

	// 解析余额
	balanceInt := new(big.Int)
	balanceInt.SetBytes(balance)
	// fmt.Printf("Balance: %s\n", balanceInt)
	return balanceInt, nil
}

func SendTransaction(transaction TransactionSt) (string, error) {
	privateKey, err := crypto.HexToECDSA(transaction.SK)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// log.Fatal("error casting public key to ECDSA")
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	chainID, err := EthClient.NetworkID(context.Background()) // 确保是BSC的Chain ID

	// 设置EIP-1559交易的参数
	// 使用types.NewTx创建EIP-1559类型的交易
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: transaction.MaxPriorityFeePerGas, // 设置合适的 MaxPriorityFeePerGas
		GasFeeCap: transaction.MaxFeePerGas,         // 设置合适的 MaxFeePerGas
		Gas:       transaction.GasLimit,             // 设置合适的 Gas Limit
		To:        transaction.ContractAddress,
		Value:     transaction.Value,
		Data:      transaction.Data,
	})

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	err = EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// TODO:
// MQ提高稳定性
// DB数据持久化
func CheckTransaction(tx string) (bool, error) {
	txHash := common.HexToHash(tx)

	var checkcount = 0

	// 轮询检查交易是否被确认
	for {
		if checkcount == 15 {
			break
		}
		_, isPending, err := EthClient.TransactionByHash(context.Background(), txHash)
		if err != nil {
			// log.Fatalf("Error getting transaction: %v", err)
			return false, err
		}

		if !isPending {
			receipt, err := EthClient.TransactionReceipt(context.Background(), txHash)
			if err != nil {
				// log.Fatalf("Error getting transaction receipt: %v", err)
				return false, err
			}

			if receipt.Status == types.ReceiptStatusSuccessful {
				// fmt.Println("Transaction confirmed!")
				return true, nil
			} else {
				// fmt.Println("Transaction failed!")
				return false, nil
			}
		}

		checkcount++
		time.Sleep(2 * time.Second) // 每2秒检查一次
	}

	return false, errors.New("transaction no found")
}
