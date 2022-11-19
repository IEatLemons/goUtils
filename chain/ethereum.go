package chain

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var ETHCli *ethclient.Client

// var EthCli *
func InitEth(Url string) {
	cli, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatal(err)
	}
	ETHCli = cli
}

func NewEth() *ethclient.Client {
	if ETHCli == nil {
		log.Fatal("ETHCli is nil")
	}
	return ETHCli
}

func GetPrivateKey(EnvFile string) (privateKey *ecdsa.PrivateKey) {
	err := godotenv.Load(EnvFile)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyStr := os.Getenv("ETHPrivateKey")
	if privateKeyStr == "" {
		privateKey = CreatePrivate()
	} else {
		privateKey = loadPrivate(privateKeyStr)
	}
	return
}

func CreatePrivate() (privateKey *ecdsa.PrivateKey) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("privateKey: %x\n", privateKey)
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// fmt.Printf("privateKeyBytes: %x\n", privateKeyBytes)

	fmt.Println("privateKey for hexadecimal", "0x"+hexutil.Encode(privateKeyBytes)[2:])
	return
}

func loadPrivate(privateKeyStr string) (privateKey *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetPublicKey(privateKey *ecdsa.PrivateKey) (publicKeyECDSA *ecdsa.PublicKey) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return
}

func EthToWei(eth *big.Float) *big.Int {
	return AmountToToken(eth, 18)
}

func WeiToEth(wei *big.Int) *big.Float {
	return TokenToAmount(wei, 18)
}

func AmountToToken(amount *big.Float, decimals int) (i *big.Int) {
	i = new(big.Int)
	multiple := new(big.Float).SetFloat64(math.Pow10(decimals))
	amount.Mul(amount, multiple).Int(i)
	return
}

func TokenToAmount(token *big.Int, decimals int) *big.Float {
	divisor := new(big.Float).SetInt(token)
	dividend := new(big.Float).SetFloat64(math.Pow10(decimals))
	return new(big.Float).Quo(divisor, dividend)
}
