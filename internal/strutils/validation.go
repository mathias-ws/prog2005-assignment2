package strutils

import (
	"assignment-2/internal/constants"
	"regexp"
	"strings"
	"time"
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
		if regexp.MustCompile(regexCheckValidDate).MatchString(strToCheck) {
			parsedTime, err := time.Parse(constants.URL_PARAMETER_WANTED_TIME_FORMAT, str)
			if err != nil {
				return false
			}

			if time.Now().After(parsedTime) {
				return true
			}
		}
	}

	return false
}
