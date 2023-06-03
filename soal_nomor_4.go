package jakpat

import (
	"fmt"
	"time"
)

var andi = 12
var beni = 14
var ahmad = 6
var startDate = time.Now().Format("2023-04-28")

func CheckMod(input int) bool {
	if (input%andi == 0) && (input%beni == 0) && (input%ahmad == 0) {
		return true
	}
	return false
}

func SearchSimilarSchedule(input string) {
	now := time.Now().Format(input + "12-31")
	fmt.Println(now)
	past, err := time.Parse(now, startDate)
	if err != nil {
		panic(err)
	}

	diff := time.Since(past)

	fmt.Println(diff.Hours())
}
