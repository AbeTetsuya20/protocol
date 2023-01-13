package rsa

import (
	"fmt"
	"github.com/AbeTetsuya20/protocol/util"
	"math"
	"math/big"
)

// 公開鍵を与えたら、秘密鍵を特定する関数
func SolveRSA(key *PublicKey) (*PrivateKey, error) {
	p, q, err := Factorize(key.N)
	if err != nil {
		return nil, err
	}

	d := GetD(p, q, key.E)
	return &PrivateKey{D: d}, nil
}

// n を 2 つの素数に素因数分解する
func Factorize(n *big.Int) (*big.Int, *big.Int, error) {
	// エラトステネスの篩で 2 ~ sqrt(n) までの素数を取得する
	primes, err := SieveOfEratosthenes(int64(math.Sqrt(float64(n.Int64()))))
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(primes)

	// 素因数分解
	for _, prime := range primes {
		modN := big.NewInt(1).Mod(n, prime)
		if modN.Cmp(big.NewInt(0)) == 0 {
			return prime, n.Div(n, prime), nil
		}
	}

	fmt.Println("素因数分解できませんでした")
	return nil, nil, nil
}

// エラトステネスの篩で 2 ~ n までの素数を取得する
func SieveOfEratosthenes(n int64) ([]*big.Int, error) {
	// 2 ~ n までの配列を作る
	numbers := make([]*big.Int, 0, n)
	for i := big.NewInt(2); i.Cmp(big.NewInt(n)) <= 0; i.Add(i, big.NewInt(1)) {
		numbers = append(numbers, big.NewInt(1).Set(i))
	}

	primes := make([]*big.Int, 0, n)

	// numbers[0] = 2
	// numbers[0] を primes に追加
	// numbers[0] の倍数を numbers から削除 (= 0 にする)
	// -> numbers[2] = 4, numbers[4] = 6, numbers[6] = 8, ...
	// -> numbers[2*j - 2] = 0
	// numbers = [0, 3, 0, 5, 0, 7...]

	for i := 0; i < len(numbers); i++ {
		// 0 だったら次のループへ
		if numbers[i].Cmp(big.NewInt(0)) == 0 {
			continue
		}

		// numbers[i] を primes に追加
		primes = append(primes, numbers[i])

		// numbers[i] の倍数を numbers から削除 (= 0 にする)
		for j := 2; j < len(numbers); j++ {

			// 2j -2 が n より大きくなったら終了
			if (i+2)*j-2 >= len(numbers) {
				break
			}

			// numTmp = (i+2) * j -2
			numTmp := big.NewInt(1).Mul(big.NewInt(int64(i+2)), big.NewInt(int64(j)))
			numTmp.Sub(numTmp, big.NewInt(2))

			// numTmp を int に変換
			numInt := numTmp.Int64()
			numbers[numInt] = big.NewInt(0)
		}
	}

	return primes, nil
}

// 2 つの数字 p, q から、逆元 d を求める
func GetD(p, q, e *big.Int) *big.Int {
	// (p-1),　(q-1)
	pMinus1 := big.NewInt(1).Sub(p, big.NewInt(1))
	qMinus1 := big.NewInt(1).Sub(q, big.NewInt(1))

	// l = lcm(p-1, q-1)
	l := util.Lcm(pMinus1, qMinus1)

	// d = e^-1 mod l
	d := big.NewInt(1).ModInverse(e, l)

	return d
}
