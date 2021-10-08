package db

import (
	"auto-garden/consts"
	"strconv"
	"unicode"
)


func IsBool(param string) bool {
	return param == "0" || param == "1"
}

func IsInt2(param string) bool {
	return isIntBySize(param, 16)
}

func IsInt8(param string) bool {
	return isIntBySize(param, 64)
}

func isIntBySize(param string, bitSize int) bool {
	if _, err := strconv.ParseInt(param, 10, bitSize); err == nil {
		return true
	}
	return false
}

func IsValidPassword(param string) bool {
	return isValidLength(param, consts.MinPasswordSize, consts.MaxPasswordSize) && hasValidPasswordChars(param)
}

func IsValidUser(param string) bool {
	return isValidLength(param, consts.MinUserSize, consts.MaxUserSize) && hasValidUserChars(param)
}

func isValidLength(param string, minLength int, maxLength int) bool {
	size := len(param)
	return size > minLength && size < maxLength
}

func hasValidPasswordChars(param string) bool {
	for _, key := range param {
		if !unicode.IsLetter(key) && !unicode.IsDigit(key) && !unicode.IsSymbol(key) {
			return false
		}
	}
	return true
}

func hasValidUserChars(param string) bool {
	for _, key := range param {
		if !unicode.IsLetter(key) && !unicode.IsDigit(key) {
			return false
		}
	}
	return true
}

