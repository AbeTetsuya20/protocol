package util

import (
	"crypto/rand"
	"math/big"
)

// bitSize の素数を生成する関数
func MakePrime(bitSize int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bitSize/2)
}

// IsPrime は n が素数かどうかを判定する関数
func IsPrime(n *big.Int) (bool, error) {
	count := 0

	for {
		if count >= 40 {
			return true, nil
		}

		a, err := rand.Int(rand.Reader, n)
		if err != nil {
			return false, err
		}

		// a^(n-1) mod n = 1 となるかチェック
		tmp := big.NewInt(1)
		nMinus1 := new(big.Int).Sub(n, big.NewInt(1))
		tmp.Exp(a, nMinus1, n)

		if tmp.Cmp(big.NewInt(1)) != 0 {
			return false, nil
		} else {
			count++
		}

	}
}

func Lcm(a, b *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(a, b), Gcd(a, b))
}

func Gcd(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	return Gcd(b, new(big.Int).Mod(a, b))
}
