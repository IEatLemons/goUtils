package sui

import (
	"log"

	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/tyler-smith/go-bip39"
)

type SuiAccounts struct {
	List []*SuiAccount
}

type SuiAccount struct {
	Index  uint64
	Signer *signer.Signer
}

func NewMnemonic() string {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatalln(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalln(err)
	}
	return mnemonic
}

func NewAccount(Mnemonic string) *SuiAccount {
	signerAccount, err := NewSignerWithMnemonic(Mnemonic, "0")
	if err != nil {
		log.Println(err.Error())
	}
	return &SuiAccount{
		Signer: signerAccount,
	}
}

func NewSignerWithMnemonic(mnemonic, index string) (*signer.Signer, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}
	key, err := signer.DeriveForPath("m/44'/784'/"+index+"'/0'/0'", seed)
	if err != nil {
		return nil, err
	}
	return signer.NewSigner(key.Key), nil
}
