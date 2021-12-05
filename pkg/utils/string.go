package utils

import "strconv"

func InvertBinaryString(s string) string {
	r := ""
	for _, v := range s {
		if v == '0' {
			r += "1"
		} else {
			r += "0"
		}
	}
	return r
}

func BinaryStringToInt(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}
