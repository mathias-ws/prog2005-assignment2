package strutils

import "testing"
import "github.com/stretchr/testify/assert"

func TestCheckIfStringIsNotEmptyValid(t *testing.T) {
	assert.Equal(t, true, CheckIfStringIsNotEmpty("this is a cool string"))
}

func TestCheckIfStringIsNotEmptySpecialChar(t *testing.T) {
	assert.Equal(t, false, CheckIfStringIsNotEmpty("this string will not work because of the !!"))
}

func TestCheckIfStringIsNotEmptyEmptyString(t *testing.T) {
	assert.Equal(t, false, CheckIfStringIsNotEmpty(""))
}

func TestCheckIfStringIsNotEmptySpace(t *testing.T) {
	assert.Equal(t, false, CheckIfStringIsNotEmpty(" "))
}

func TestCheckIfStringIsNotEmptySpaceAndNumber(t *testing.T) {
	assert.Equal(t, false, CheckIfStringIsNotEmpty("number 2"))
}

func TestCheckIfStringIsNotEmptyCamelCase(t *testing.T) {
	assert.Equal(t, true, CheckIfStringIsNotEmpty("superCoolString"))
}

func TestCheckIfValidDateFormatValid(t *testing.T) {
	assert.Equal(t, true, CheckIfValidDateFormat("2021-12-22"))
}

func TestCheckIfValidDateFormatMissingDash(t *testing.T) {
	assert.Equal(t, false, CheckIfValidDateFormat("202212-22"))
}

func TestCheckIfValidDateFormatValidDate(t *testing.T) {
	assert.Equal(t, true, CheckIfValidDateFormat("2020-03-22"))
}

func TestCheckIfValidDateFormatZeroMonth(t *testing.T) {
	assert.Equal(t, false, CheckIfValidDateFormat("2022-2-22"))
}

func TestCheckIfValidDateFormatZeroDate(t *testing.T) {
	assert.Equal(t, false, CheckIfValidDateFormat("2022-12-2"))
}

func TestCheckIfValidDateFormatValidDate2(t *testing.T) {
	assert.Equal(t, true, CheckIfValidDateFormat("2021-12-02"))
}

func TestCheckIfValidDateFormatDateInFuture(t *testing.T) {
	assert.Equal(t, false, CheckIfValidDateFormat("2030-12-02"))
}
