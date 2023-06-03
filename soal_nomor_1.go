package jakpat

import (
	"strings"
)

func ValidatePhoneNumber(value string) string {
	newString := strings.ReplaceAll(value, "-", "")
	arrStr := []rune(newString)
	var strNewPhoneNumber string
	flag := false
	for i := 0; i < len(arrStr); i++ {
		tempIndex := i
		if flag {
			tempIndex += 1
		}
		if string(arrStr[i]) != "0" && i == 0 {
			strNewPhoneNumber += "0"
			flag = true
		} else if (tempIndex%4 == 0) && tempIndex != 0 {
			strNewPhoneNumber += "-"
		}

		strNewPhoneNumber += string(arrStr[i])
	}

	return strNewPhoneNumber
}
