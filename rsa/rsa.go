package rsa

import (
	"github.com/AbeTetsuya20/protocol/util"
	"math/big"
)

// 公開鍵の構造体
type PublicKey struct {
	E *big.Int
	N *big.Int
}

// 秘密鍵の構造体
type PrivateKey struct {
	D *big.Int
}

// RSA 暗号の鍵を生成する
func MakeKeys(bitSize int) (*PublicKey, *PrivateKey, error) {

	// 2 つの bitSize/2 bit の素数を生成する
	p, err := util.MakePrime(bitSize / 2)
	if err != nil {
		return nil, nil, err
	}

	q, err := util.MakePrime(bitSize / 2)
	if err != nil {
		return nil, nil, err
	}

	// n = p * q
	n := big.NewInt(1)
	n.Mul(p, q)

	// l = lcm(p-1, q-1)
	tmpP := big.NewInt(1)
	tmpP.Sub(p, big.NewInt(1))
	tmpQ := big.NewInt(1)
	tmpQ.Sub(q, big.NewInt(1))
	l := util.Lcm(tmpP, tmpQ)

	// e は l と互いに素な整数
	e, err := util.MakePrime(bitSize / 2)
	if err != nil {
		return nil, nil, err
	}

	// d は e と l の逆元
	d := big.NewInt(1)
	d.ModInverse(e, l)

	// 秘密鍵と公開鍵を生成
	privateKey := PublicKey{e, n}
	publicKey := PrivateKey{d}

	return &privateKey, &publicKey, nil
}

func RSAEncrypt(message, n, e *big.Int) *big.Int {
	return new(big.Int).Exp(message, e, n)
}

func RSADecrypt(message, n, d *big.Int) *big.Int {
	return new(big.Int).Exp(message, d, n)
}
