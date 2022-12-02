package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeyPair() (publicKey string, privateKey string, err error) {
	var key *ecdsa.PrivateKey

	if key, err = crypto.GenerateKey(); err != nil {
		return
	}
	publicKey = hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey))
	privateKey = hex.EncodeToString(key.D.Bytes())
	return
}
