package evm

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type OpContract struct {
	chainID *big.Int
	client  *ethclient.Client
	maxGas  *big.Int
	minGas  *big.Int
}

func NewOpContract(
	chainID,
	maxGas,
	minGas *big.Int,
	rpcPath string,
) *OpContract {
	client, err := ethclient.Dial(rpcPath)
	if err != nil {
		log.Fatalln(err)
	}
	return &OpContract{
		chainID: chainID,
		client:  client,
		maxGas:  maxGas,
		minGas:  minGas,
	}
}

func (c *OpContract) ChainID() *big.Int {
	return c.chainID
}

func (c *OpContract) Client() *ethclient.Client {
	return c.client
}

func (c *OpContract) GasPrice() *big.Int {
	return c.chainID
}

func (c *OpContract) Gas() (gasPrice *big.Int) {
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if gasPrice.Cmp(c.maxGas) > 0 {
		gasPrice = c.maxGas
	}
	if gasPrice.Cmp(c.minGas) < 0 {
		gasPrice = c.minGas
	}
	return
}

func (c *OpContract) ListenGas() (gasPrice *big.Int) {
RESET:
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if gasPrice.Cmp(c.maxGas) > 0 {
		log.Println("The current gas is", TokenToAmount(gasPrice, 9).String(), "Gwei, But MaxGas is", TokenToAmount(c.maxGas, 9).String(), "Gwei")
		time.Sleep(time.Second * 10)
		goto RESET
	}
	if gasPrice.Cmp(c.minGas) < 0 {
		gasPrice = c.minGas
	}
	return
}
