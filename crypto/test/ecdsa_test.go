package crypto_test

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

// Simple signing process
func TestSigning(T *testing.T) {
	privateKey := getPrivate()
	publicKeyECDSA, publicKeyBytes := seePublic(privateKey)
	seeAddress(publicKeyECDSA, publicKeyBytes)

	msg := "hellp, world"
	hash := sha256.Sum256([]byte(msg))
	fmt.Printf("tx hash: %x\n", hash[:])
	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}

func getPrivate() (privateKey *ecdsa.PrivateKey) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	privateKeyStr := os.Getenv("ETHPrivateKey")
	if privateKeyStr == "" {
		privateKey = createPrivate()
	} else {
		privateKey = loadPrivate(privateKeyStr)
	}
	return
}

func createPrivate() (privateKey *ecdsa.PrivateKey) {
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

func seePublic(privateKey *ecdsa.PrivateKey) (publicKeyECDSA *ecdsa.PublicKey, publicKeyBytes []byte) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// fmt.Println("publicKeyECDSA", publicKeyECDSA)

	publicKeyBytes = crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKey for hexadecimal", "0x"+hexutil.Encode(publicKeyBytes)[4:])
	return
}

func seeAddress(publicKeyECDSA *ecdsa.PublicKey, publicKeyBytes []byte) {
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("PubkeyToAddress Address is", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Keccak256 Address is", hexutil.Encode(hash.Sum(nil)[12:]))
}
