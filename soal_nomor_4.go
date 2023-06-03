package jakpat

import (
	"time"
)

var andi = 12
var beni = 14
var ahmad = 6
var startDate = time.Date(2023, 4, 28, 0, 0, 0, 0, time.UTC)

func CheckMod(input int) bool {
	if (input%andi == 0) && (input%beni == 0) && (input%ahmad == 0) {
		return true
	}
	return false
}

func SearchSimilarSchedule(input int) []string {
	now := time.Date(input, 12, 31, 0, 0, 0, 0, time.UTC)
	var dayTotal int
	if input == 2023 {
		diff := now.Sub(startDate)
		dayTotal = int(diff.Hours()) / 24
	} else if input > 2023 {
		dayTotal = 365
	} else {
		return []string{}
	}

	var arrDate []time.Time
	for i := 0; i < dayTotal; i++ {
		if CheckMod(i) {
			hours := i * 24
			startDate.Add(time.Duration(hours))
			arrDate = append(arrDate, startDate)
		}
	}
	return []string{}
}
