package utilities

// RegexCheckValidDate regex expression for checking if string is valid date format.
// Inspired by: https://www.golangprograms.com/regular-expression-to-validate-the-date-format-in-dd-mm-yyyy.html
const regexCheckValidDate string = "((19|20)\\d\\d)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])"

// RegexCheckValidString regex expression that checks that a string contains small or capital letters.
const regexCheckValidString string = "^[a-zA-Z]+$"
