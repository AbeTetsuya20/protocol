package main

import (
	"fmt"
	"github.com/AbeTetsuya20/protocol/advance"
	"github.com/AbeTetsuya20/protocol/rsa"
	"math/big"
)

func main() {

	message := big.NewInt(100)

	publicKey, privateKey, err := rsa.MakeKeys(50)
	if err != nil {
		panic(err)
	}
	messageEnc := rsa.RSAEncrypt(message, publicKey.N, publicKey.E)
	messageEncDec := rsa.RSADecrypt(messageEnc, publicKey.N, privateKey.D)

	fmt.Printf("message: %v \n", message)
	fmt.Printf("暗号化した message: %v \n", messageEnc)
	fmt.Printf("暗号化した message を復号化した message: %v \n", messageEncDec)

	fmt.Printf("公開鍵: %+v \n", publicKey)
	fmt.Printf("秘密鍵: %+v \n", privateKey)

	privateKeySolve, err := advance.SolveRSA(publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("特定した秘密鍵: %+v \n", privateKeySolve)
}
