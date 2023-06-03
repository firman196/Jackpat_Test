package jakpat

import "strconv"

var totalPages = 231

func SearchPage(page int) string {
	var result int = 0
	half := (totalPages + 1) / 2
	if page <= half {
		for i := 0; i <= page; i++ {
			if i%2 == 0 && i != 0 {
				result += 1
			}
		}
		return "From first page flip " + strconv.Itoa(result) + "x"
	} else {
		for i := totalPages; i > page; i-- {
			if i%3 == 0 && i != totalPages {
				result += 1
			}
		}
		return "From last page flip " + strconv.Itoa(result) + "x"
	}

}
