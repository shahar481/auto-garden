package db

import "strconv"

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
