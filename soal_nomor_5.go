package jakpat

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Encription(input string) string {
	str := strings.ReplaceAll(input, " ", "")
	arrStr := []rune(str)
	total := len(str)
	akar := math.Sqrt(float64(total))
	strFloat := fmt.Sprintf("%v", akar)
	row, _ := strconv.Atoi(string(strFloat[0]))
	f, _ := strconv.ParseFloat(string(strFloat[0]), 64)
	col := row
	if akar > f {
		col += 1
	}

	var arrCol [][]string
	var tempString []string
	for i := 0; i < len(str); i++ {
		if len(arrCol) <= row {
			if i == len(str)-1 {
				tempString = append(tempString, string(arrStr[i]))
			}

			if col > len(tempString) && i != len(str)-1 {
				tempString = append(tempString, string(arrStr[i]))
			} else {
				arrCol = append(arrCol, tempString)
				tempString = []string{}
				tempString = append(tempString, string(arrStr[i]))
			}
		}

	}

	var result string
	for l := 0; l <= col; l++ {
		for r := 0; r < len(arrCol); r++ {
			if len(arrCol[r]) > l {
				result += string(arrCol[r][l])
			}

		}
		if l < col-1 {
			result += " "
		}
	}

	return result

}
