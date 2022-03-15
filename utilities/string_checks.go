package utilities

import (
	"regexp"
	"strings"
)

// CheckIfStringIsNotEmpty checks if a string contains only letters.
func CheckIfStringIsNotEmpty(str string) bool {
	//Removes whitespaces
	strToCheck := strings.ReplaceAll(str, " ", "")

	if strToCheck != "" {
		return regexp.MustCompile(regexCheckValidString).MatchString(strToCheck)
	}
	return false
}

// CheckIfValidDateFormat takes a string and checks if it is a valid date format.
func CheckIfValidDateFormat(str string) bool {
	//Removes whitespaces
	strToCheck := strings.ReplaceAll(str, " ", "")

	if strToCheck != "" {
		return regexp.MustCompile(regexCheckValidDate).MatchString(strToCheck)
	}

	return false
}
