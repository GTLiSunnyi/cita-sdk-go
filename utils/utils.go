package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/tjfoc/gmsm/sm3"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
)

func ParseAddress(str string) ([]byte, error) {
	address, err := ParseData(str)
	if err != nil {
		return nil, err
	}
	if len(address) == types.AddressSize {
		return address, nil
	} else {
		return nil, errors.New("地址长度错误")
	}
}

func ParseValue(str string) ([]byte, error) {
	str = remove0x(str)

	// 前置补0
	str, err := PreSupplyZero(str, 64)
	if err != nil {
		return nil, err
	}

	res, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ParseData(str string) ([]byte, error) {
	return hex.DecodeString(remove0x(str))
}

func remove0x(str string) string {
	return strings.Replace(str, "0x", "", 1)
}

/*
PreSupplyZero
@Description: 字符串前置补0
@param str: 需要操作的字符串
@param length: 结果字符串的长度
@return string
*/
func PreSupplyZero(str string, length int) (string, error) {
	if len(str) > length || length <= 0 {
		return "", errors.New("length 错误, 或者字符串长度错误")
	}

	return fmt.Sprintf("%0*s", length, str), nil
}

func Sm3Hash(data []byte) []byte {
	return sm3.New().Sum(data)
}
