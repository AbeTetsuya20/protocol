package main

import (
	"fmt"
	"github.com/AbeTetsuya20/protocol/elgamal"
	"math/big"
	"time"
)

func main() {

	start := time.Now()

	// RSA 暗号
	message := big.NewInt(1000)

	publicKey, privateKey, err := elgamal.MakeKeys(2048)
	if err != nil {
		panic(err)
	}

	// 暗号化
	messageEnc1, messageEnc2, err := elgamal.Encrypt(message, publicKey)
	if err != nil {
		panic(err)
	}

	// 復号化
	messageEncDec := elgamal.Decrypt(messageEnc1, messageEnc2, publicKey, privateKey)

	time := time.Since(start).Milliseconds()

	fmt.Printf("公開鍵 \n")
	fmt.Printf("P: %v \n", publicKey.P)
	fmt.Printf("G: %v \n", publicKey.G)
	fmt.Printf("Y: %v \n", publicKey.Y)
	fmt.Printf("秘密鍵 \n")
	fmt.Printf("X: %v \n", privateKey.X)

	fmt.Println("")
	fmt.Printf("平文　   : %s \n", message)
	fmt.Printf("暗号文(1): %s \n", messageEnc1)
	fmt.Printf("暗号文(2): %s \n", messageEnc2)
	fmt.Printf("復号文   : %s \n", messageEncDec)

	fmt.Println("")
	fmt.Printf("処理時間: %d ms \n", time)
}
