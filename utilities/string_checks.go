package utilities

import (
	"assignment-2/constants"
	"regexp"
	"strings"
)

// CheckIfStringIsNotEmpty checks if a string contains only letters.
func CheckIfStringIsNotEmpty(str string) bool {
	//Removes whitespaces
	strToCheck := strings.ReplaceAll(str, " ", "")

	if strToCheck != "" {
		return regexp.MustCompile(constants.REGEX_CHECK_VALID_STRING).MatchString(strToCheck)
	}
	return false
}

// CheckIfValidDateFormat takes a string and checks if it is a valid date format.
func CheckIfValidDateFormat(str string) bool {
	//Removes whitespaces
	strToCheck := strings.ReplaceAll(str, " ", "")

	if strToCheck != "" {
		return regexp.MustCompile(constants.REGEX_CHECK_VALID_DATE).MatchString(strToCheck)
	}

	return false
}
