package country

import (
	"assignment-2/web_client"
	"strings"
)

// buildSearchUrl builds a search url for the country api.
func buildSearchUrl(cca3Code string) string {
	urlToSearch := strings.Builder{}
	urlToSearch.WriteString(countryApiUrl)
	urlToSearch.WriteString(countryApiCodeParam)
	urlToSearch.WriteString(cca3Code)
	urlToSearch.WriteString(countryApiFieldParam)

	return urlToSearch.String()
}

// GetCountryNameFromCca3 converts the inputted cca3 code into a full country name.
func GetCountryNameFromCca3(cca3Code string) (string, error) {
	url := buildSearchUrl(cca3Code)

	response, errResponse := web_client.GetRequest(url)

	if errResponse != nil {
		return "", errResponse
	}

	country := decodeCountryInfo(response)

	return country.Name.Common, nil
}
