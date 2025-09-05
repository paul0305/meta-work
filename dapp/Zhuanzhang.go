package dapp

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
*
*代币转账
连接到本地 anvil 网络成功
发送方地址: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
当前 nonce: 0
接收方地址: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
方法ID: 0xa9059cbb
接收方地址编码: 0x00000000000000000000000070997970c51812dc3a010c7d01b50e0d17dc79c8
转账金额编码: 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
交易已发送，交易哈希: 0xeda28dde2f03479f38c0a18ad9bf7c54a786411b62b80abe657bd62016bc2916
*/
func Transport() {
	// 连接到本地 anvil 网络
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("连接到本地 anvil 网络成功")

	// 使用 anvil 提供的测试账户私钥
	// 这是 anvil 默认的第一个账户私钥
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *crypto.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("发送方地址: %s\n", fromAddress.Hex())

	// 获取 nonce 值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("当前 nonce: %d\n", nonce)

	// 设置交易参数
	value := big.NewInt(0) // 转账金额为 0 wei (ERC20 代币转账)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 接收方地址 (使用 anvil 默认的第二个账户)
	//toAddress := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	toAddress := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	fmt.Printf("接收方地址: %s\n", toAddress.Hex())

	// 合约地址 (假设已部署 ERC20 合约)
	//tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// 构造 transfer 函数调用数据
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("方法ID: %s\n", hexutil.Encode(methodID)) // 0xa9059cbb

	// 准备参数
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("接收方地址编码: %s\n", hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("转账金额编码: %s\n", hexutil.Encode(paddedAmount))

	// 构造调用数据
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 估算 gas
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &tokenAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("估算的 gas: %d\n", gasLimit)

	// 创建交易
	//tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	})

	// 获取链ID并签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送，交易哈希: %s\n", signedTx.Hash().Hex())
}

func method2() {
	//1.连接到 本地 anvil 网络
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接到本地 anvil 网络成功")
	//2.使用 anvil 默认的测试账户私钥
	//默认的第一个账户私钥：ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
	//在发送交易前，使用私钥创建一个ECDSA椭圆曲线密钥对
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80") // Use the private key here
	if err != nil {
		log.Fatal(err)
	}
	//3.通过私钥获取公钥和地址
	publicKey := privateKey.Public()
	publicEcdsaKey := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicEcdsaKey)
	fmt.Printf("发送方地址: %s\n", fromAddress.Hex())
	//4.获取nonce值
	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("当前 nonce: %d\n", nonceAt)
	//5.设置交易参数
	value := big.NewInt(0)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//接收方地址(使用 anvil 默认的第二个账户)
	toAddress := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	fmt.Printf("接收方地址: %s\n", toAddress.Hex())
	//6构造 transfer 函数调用数据 两个参数 address, uint256 转账地址，转账金额
	transferFnSignature := []byte("transfer(address,uint256)")
	keccak256 := sha3.NewLegacyKeccak256() //创建 Keccak256哈希实例
	keccak256.Write(transferFnSignature)   //向哈希实例写入 transfer 函数签名
	methodId := keccak256.Sum(nil)[:4]     //计算函数值，并且取前4位
	fmt.Printf("方法 id: %s\n", hexutil.Encode(methodId))
	//准备参数
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("接收方地址编码: %s\n", hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("转账金额编码: %s\n", hexutil.Encode(paddedAmount))
	//构造调用函数
	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	//7.创建交易
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonceAt,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//8 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//9 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("交易已发送，交易哈希: %s\n", signedTx.Hash().Hex())
}
