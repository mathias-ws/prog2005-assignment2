package country

var (
	countryApiUrl       = "https://restcountries.com/v3.1/alpha/"
	countryDbCollection = "cca3tocountry"
)

const countryApiCodeParam string = "?codes="
const countryApiFieldParam string = "&fields=name"

// SetTestCollection sets the test collection.
func SetTestCollection() {
	countryDbCollection = "testcca3tocountry"
}

// SetTestUrlCountry sets the test url for the country api.
func SetTestUrlCountry(testUrl string) {
	countryApiUrl = testUrl
}
