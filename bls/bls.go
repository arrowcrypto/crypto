package bls

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bn256"
)

// PublicKey is the BLS public key, i.e. a point on curve G2
type PublicKey struct {
	gx *bn256.G2
}

// Private key is a scalar
type PrivateKey struct {
	PublicKey
	x *big.Int
}

// Signature is a point on G1
type Signature struct {
	sigma *bn256.G1
}

func (privKey *PrivateKey) GetPublicKey() *PublicKey {
	var pubKey PublicKey
	pubKey.gx = privKey.gx
	return &pubKey
}

func GenerateKeyPair() (*PrivateKey, error) {
	var privKey PrivateKey
	var err error
	privKey.x, privKey.gx, err = bn256.RandomG2(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &privKey, nil
}

func Sign(privKey *PrivateKey, msg []byte) *Signature {

	sigma := bn256.HashToG1Point(msg)
	sigma = sigma.ScalarMult(sigma, privKey.x)
	return &Signature{sigma: sigma}
}
