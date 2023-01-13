package elgamal

import (
	"crypto/rand"
	"github.com/AbeTetsuya20/protocol/util"
	"math/big"
)

// 公開鍵の構造体
type PublicKey struct {
	P *big.Int
	G *big.Int
	Y *big.Int
}

// 秘密鍵の構造体
type PrivateKey struct {
	X *big.Int
}

// Elgamal 暗号の鍵を生成する
func MakeKeys(bitSize int) (*PublicKey, *PrivateKey, error) {

	// bitSize bit の素数 q を生成する
	// p = 2q + 1 となる p が素数になるかチェックする

	p := big.NewInt(1)
	q := big.NewInt(1)

	for {
		var err error
		q, err = util.MakePrime(bitSize)
		if err != nil {
			return nil, nil, err
		}

		p = big.NewInt(1)
		p.Mul(q, big.NewInt(2))
		p.Add(p, big.NewInt(1))

		isPrimeBool, err := util.IsPrime(p)
		if err != nil {
			return nil, nil, err
		}

		if isPrimeBool {
			break
		}

	}

	// g　は位数が p-1
	// g　をまずランダムに選ぶ
	g := big.NewInt(1)
	for {
		var err error
		g, err = rand.Int(rand.Reader, p)
		if err != nil {
			return nil, nil, err
		}

		// g^q mod p = 1 なら g は位数が p-1
		tmp := big.NewInt(1)
		tmp.Exp(g, q, p)
		if tmp.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	// x はランダムに選ぶ
	x, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, nil, err
	}

	// y = g^x mod p
	y := big.NewInt(1)
	y.Exp(g, x, p)

	// 秘密鍵と公開鍵を生成
	privateKey := PrivateKey{x}
	publicKey := PublicKey{p, g, y}

	return &publicKey, &privateKey, nil
}

// 暗号化
func Encrypt(message *big.Int, publicKey *PublicKey) (*big.Int, *big.Int, error) {
	// r をランダムに選ぶ
	r, err := rand.Int(rand.Reader, publicKey.P)
	if err != nil {
		return nil, nil, err
	}

	// c1 = g^r mod p
	// c2 = m * y^r mod p
	c1 := big.NewInt(1)
	c1.Exp(publicKey.G, r, publicKey.P)
	c2 := big.NewInt(1)
	c2.Exp(publicKey.Y, r, publicKey.P)
	c2.Mul(c2, message)
	c2.Mod(c2, publicKey.P)

	return c1, c2, nil
}

// 復号化
func Decrypt(c1 *big.Int, c2 *big.Int, publicKey *PublicKey, privateKey *PrivateKey) *big.Int {
	// m = c2 * c1^x mod p
	m := big.NewInt(1)
	m.Exp(c1, privateKey.X, publicKey.P)
	m.ModInverse(m, publicKey.P)
	m.Mul(m, c2)
	m.Mod(m, publicKey.P)

	return m
}
