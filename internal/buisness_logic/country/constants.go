package country

var countryApiUrl = "https://restcountries.com/v3.1/alpha/"

const countryApiCodeParam string = "?codes="
const countryApiFieldParam string = "&fields=name"

const countryDbCollection string = "cca3tocountry"

// SetTestUrlCountry sets the test url for the country api.
func SetTestUrlCountry(testUrl string) {
	countryApiUrl = testUrl
}
