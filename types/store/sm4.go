package store

import (
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm4"

	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type SM4 struct{}

// CFB模式加密
func (SM4) Encrypt(text string, password string) (string, error) {
	key := utils.Sm3Hash([]byte(password))[:16]
	res, err := sm4.Sm4CFB(key, []byte(text), true)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(res), nil
}

// CFB模式解密
func (SM4) Decrypt(cryptoText string, password string) (string, error) {
	text, err := hex.DecodeString(cryptoText[2:])
	if err != nil {
		return "", err
	}

	key := utils.Sm3Hash([]byte(password))[:16]
	res, err := sm4.Sm4CFB(key, text, false)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
