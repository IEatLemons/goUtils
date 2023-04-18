package evm

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SamrtContract struct {
	Client          *ethclient.Client
	ABI             abi.ABI
	ContractAddress common.Address
}

func NewSC(rpc, abiJson, ContractAddress string) (*SamrtContract, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	abi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, err
	}
	return &SamrtContract{
		Client:          client,
		ContractAddress: common.HexToAddress(ContractAddress),
		ABI:             abi,
	}, nil
}

func (SC *SamrtContract) Call(funcName string, args ...interface{}) (returnValue []interface{}, err error) {
	callData, err := SC.ABI.Pack(funcName, args...)
	if err != nil {
		return
	}
	result, err := SC.Client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &SC.ContractAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return
	}

	returnValue, err = SC.ABI.Unpack(funcName, result)
	return
}
