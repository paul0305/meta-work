package dapp_task1

import (
	"context"
	"crypto/ecdsa"
	"dapp-task1/counter"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	contractAdd = "0x97698ef2A759995bea0E42e6A35dcD27f5204Ab9"
	spoliaKey   = "https://sepolia.infura.io/v3/df757289876e4151b63ac67474868958"
)

/*
*
合约的调用
*/
func exeAdd() {
	client, err := ethclient.Dial(spoliaKey)
	if err != nil {
		panic(err)
	}
	//创建合约实例
	counterContract, err := counter.NewCounter(common.HexToAddress(contractAdd), client)
	if err != nil {
		panic(err)
	}
	// 需要提供有效的私钥才能发送交易
	fmt.Println("要调用 addNb() 方法，请提供有效的私钥来发送交易")
	privatekey, err := crypto.HexToECDSA("1b38fe1fecabc43b98fe2566cb4f4a58d797288f598be44c6d6dd63fc8806f4d")

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		log.Fatal(err)
	}
	//获取gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//创建一个合约交易
	auth, err := bind.NewKeyedTransactorWithChainID(privatekey, big.NewInt(11155111)) // Sepolia chain ID
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	//通过合约实例调用合约方法
	tx, err := counterContract.AddNb(auth)
	if err != nil {
		log.Fatal("Failed to execute AddNb:", err)
	}
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())
	fmt.Println("等待交易确认...")
	fmt.Printf("交易数据:%s\n", string(tx.Data()))
	fmt.Printf("交易发送方地址:%s\n", tx.Cost())
	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("key"))
	copy(value[:], []byte("value"))
}
