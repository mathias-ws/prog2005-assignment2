package utilities

import "testing"
import "github.com/stretchr/testify/assert"

func TestCheckIfStringIsNotEmpty(t *testing.T) {
	assert.Equal(t, true, CheckIfStringIsNotEmpty("this is a cool string"))
	assert.Equal(t, false, CheckIfStringIsNotEmpty("this string will not work because of the !!"))
	assert.Equal(t, false, CheckIfStringIsNotEmpty(""))
	assert.Equal(t, false, CheckIfStringIsNotEmpty(" "))
	assert.Equal(t, false, CheckIfStringIsNotEmpty("number 2"))
	assert.Equal(t, true, CheckIfStringIsNotEmpty("superCoolString"))
}

func TestCheckIfValidDateFormat(t *testing.T) {
	assert.Equal(t, true, CheckIfValidDateFormat("2022-12-22"))
	assert.Equal(t, false, CheckIfValidDateFormat("202212-22"))
	assert.Equal(t, true, CheckIfValidDateFormat("2020-03-22"))
	assert.Equal(t, false, CheckIfValidDateFormat("2022-2-22"))
	assert.Equal(t, false, CheckIfValidDateFormat("2022-12-2"))
	assert.Equal(t, true, CheckIfValidDateFormat("2022-12-02"))
}
