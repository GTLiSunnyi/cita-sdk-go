package utils

import (
	"encoding/hex"
	"strings"

	"github.com/tjfoc/gmsm/sm3"
)

func ParseData(str string) ([]byte, error) {
	return hex.DecodeString(strings.Split(str, "0x")[1])
}

func Sm3Hash(data []byte) []byte {
	return sm3.New().Sum(data)
}
