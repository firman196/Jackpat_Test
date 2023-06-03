package jakpat

import (
	"regexp"
	"strings"
)

func SearchChat(substr string) []string {

	chatRoom := map[string][]map[string]string{
		"backend": []map[string]string{
			map[string]string{
				"user":    "Ari",
				"message": "Halo",
			},
			map[string]string{
				"user":    "Ariyanto",
				"message": "Halo Juga",
			},
		},
		"frontend": []map[string]string{
			map[string]string{
				"user":    "suryanto",
				"message": "Lorem Ipsum",
			},
			map[string]string{
				"user":    "Arindo",
				"message": "Hello",
			},
		},
	}
	var response []string

	for index, values := range chatRoom {
		getValuesString := GetValue(values)
		str := strings.Join(getValuesString, ",")
		matched, _ := regexp.MatchString(strings.ToUpper(substr), strings.ToUpper(str))
		if matched {
			response = append(response, index)
		}

	}
	return response
}

func GetValue(params []map[string]string) []string {
	var tempString []string
	for _, values := range params {
		for _, value := range values {
			tempString = append(tempString, value)
		}
	}
	return tempString
}
