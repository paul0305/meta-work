package dapp_task1

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var sepoliaUrl = "https://sepolia.infura.io/v3/df757289876e4151b63ac67474868958"

/**
 * 搜索区块
 */
func searhBloc() error {
	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		return err
	}
	blockNum := big.NewInt(9124471)
	//获取区块头
	header, err := client.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		return err
	}
	fmt.Printf("获取区块的时间戳:%d\n", header.Time)
	fmt.Printf("获取区块的哈希值:%s\n", header.Hash().Hex())
	fmt.Printf("获取区块的难度值:%d\n", header.Difficulty.Uint64())
	fmt.Printf("获取区块的GasLimit:%d\n", header.GasLimit)
	return nil
	/**
	*获取区块的时间戳:1756901472
	获取区块的哈希值:0xa190ab57d101af8c8c1386a9220164c022fb26499268478f2e2f33a1b9631488
	获取区块的难度值:0
	获取区块的GasLimit:59940949
	*/
}
