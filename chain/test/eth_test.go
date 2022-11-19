package chain_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/IEatLemons/goUtils/chain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestBase(t *testing.T) {

	chain.InitEth("http://127.0.0.1:8545")
	eth_getBalance(chain.ETHCli)
}

func eth_getBalance(client *ethclient.Client) {
	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}

func TestCreate(t *testing.T) {
	_ = chain.CreatePrivate()
}
