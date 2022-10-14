package sm2

import (
	"crypto/rand"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type SM2KeyPair struct {
	types.KeyType
	Address    []byte
	PrivateKey *sm2.PrivateKey
	PublicKey  *sm2.PublicKey
}

type sm2Signature struct {
	R, S *big.Int
}

func NewKeyPair() (types.KeyPair, error) {
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	publicKey := &privateKey.PublicKey

	address := utils.Sm3Hash(append(publicKey.X.Bytes(), publicKey.Y.Bytes()...))[32-types.AddressSize:]

	return SM2KeyPair{
		KeyType:    types.Sm2Type,
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func (keypair SM2KeyPair) GetPrivateKey() string {
	return "0x" + hex.EncodeToString(keypair.PrivateKey.D.Bytes())
}

func (keypair SM2KeyPair) GetPublicKey() string {
	return "0x" + hex.EncodeToString(append(keypair.PublicKey.X.Bytes(), keypair.PublicKey.Y.Bytes()...))
}

func (keypair SM2KeyPair) GetAddress() []byte {
	return keypair.Address
}

func (keypair SM2KeyPair) GetAddressString() string {
	return "0x" + fmt.Sprintf("%x", keypair.Address)
}

func (keypair SM2KeyPair) Type() types.KeyType {
	return keypair.KeyType
}

func (keypair SM2KeyPair) Sign(msg []byte) ([]byte, error) {
	sigBytes, err := keypair.PrivateKey.Sign(nil, msg, nil)
	if err != nil {
		return nil, err
	}

	var sign sm2Signature
	_, err = asn1.Unmarshal(sigBytes, &sign)
	if err != nil {
		return nil, err
	}

	pkBytes := append(keypair.PublicKey.X.Bytes(), keypair.PublicKey.Y.Bytes()...)

	return append(append(sign.R.Bytes(), sign.S.Bytes()...), pkBytes...), nil
}

func (keypair SM2KeyPair) Verify(msg []byte, sig []byte) bool {
	return keypair.PublicKey.Verify(msg, sig)
}

func ImportKeyPair(str string) (types.KeyPair, error) {
	privKey, err := hex.DecodeString(str[2:])
	if err != nil {
		return nil, err
	}

	k := new(big.Int).SetBytes(privKey)
	c := sm2.P256Sm2()
	privateKey := new(sm2.PrivateKey)
	privateKey.PublicKey.Curve = c
	privateKey.D = k
	privateKey.PublicKey.X, privateKey.PublicKey.Y = c.ScalarBaseMult(k.Bytes())

	publicKey := &privateKey.PublicKey
	address := utils.Sm3Hash(append(publicKey.X.Bytes(), publicKey.Y.Bytes()...))[32-types.AddressSize:]

	return SM2KeyPair{
		KeyType:    types.Sm2Type,
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}
