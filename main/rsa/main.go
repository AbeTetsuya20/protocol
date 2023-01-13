package main

import (
	"fmt"
	"github.com/AbeTetsuya20/protocol/rsa"
	"math/big"
)

func main() {

	// RSA 暗号
	message := big.NewInt(100)

	publicKey, privateKey, err := rsa.MakeKeys(100)
	if err != nil {
		panic(err)
	}

	// 暗号化
	messageEnc := rsa.RSAEncrypt(message, publicKey)

	// 復号化
	messageEncDec := rsa.RSADecrypt(messageEnc, publicKey, privateKey)

	fmt.Printf("公開鍵 \n")
	fmt.Printf("E: %v \n", publicKey.E)
	fmt.Printf("N: %v \n", publicKey.N)
	fmt.Printf("秘密鍵 \n")
	fmt.Printf("D: %v \n", privateKey.D)

	fmt.Println("")
	fmt.Printf("平文　: %s \n", message)
	fmt.Printf("暗号文: %s \n", messageEnc)
	fmt.Printf("復号文: %s \n", messageEncDec)

	//fmt.Println("解読中")
	//
	//privKey, err := advance.SolveRSA(publicKey)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(privKey)

	//rsa.TimeRSA()
	rsa.SolveRSATime()
}
