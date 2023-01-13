package rsa

import (
	"fmt"
	"math/big"
	"time"
)

// RSA 暗号の鍵生成から暗号化までの時間を計測
func TimeRSA() {
	message := big.NewInt(100)

	// 30 bit ~ 1024 bit までの鍵を生成して、時間を計測する
	timeList := make([]float64, 0)
	bitList := make([]int, 0)

	for i := 30; i <= 2048; i += 20 {
		start := time.Now()

		for j := 0; j < 10; j++ {
			public, private, _ := MakeKeys(i)
			messageEnc := RSAEncrypt(message, public)
			_ = RSADecrypt(messageEnc, public, private)
		}

		end := time.Now()

		timeRSA := end.Sub(start).Seconds()
		timeList = append(timeList, timeRSA/3)
		bitList = append(bitList, i)
	}

	fmt.Printf("%#v \n", timeList)
	fmt.Println(bitList)
}

// RSA 暗号を解読するまでの時間を計測
func SolveRSATime() {
	// 30 bit ~ 1024 bit までの鍵を生成して、時間を計測する
	timeList := make([]float64, 0)
	bitList := make([]int, 0)

	// RSA 暗号を i bit で作成
	for i := 30; i < 106; i++ {
		fmt.Println(i, "bit で作成")
		public, _, _ := MakeKeys(i)

		// RSA 暗号を解読するまでの時間を計測
		start := time.Now()
		_, _ = SolveRSA(public)
		end := time.Now()
		timeRSA := end.Sub(start).Seconds()

		timeList = append(timeList, timeRSA)
		bitList = append(bitList, i)

		//if timeRSA >= 15 {
		//	break
		//}
	}

	fmt.Printf("%#v \n", timeList)
	fmt.Printf("%#v \n", bitList)
}
