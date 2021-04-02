package locale

import (
	"strings"
)

// Adress verify if string is valid in context
func Adress(loc string) string {
	localeValids := []string{"street", "avenue", "road"}
	firstWordAdress := strings.Split(strings.ToLower(loc), " ")[0]

	localeValidStatus := false

	for _, value := range localeValids {
		if value == firstWordAdress {
			localeValidStatus = true
		}
	}

	if localeValidStatus {
		return strings.Title(firstWordAdress)
	}

	return "locale invalid."
}
