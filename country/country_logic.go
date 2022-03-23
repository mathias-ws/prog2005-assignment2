package country

import (
	"assignment-2/database"
	"assignment-2/web_client"
	"strings"
	"time"
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
	countryObtainedFromDb := database.GetFromDatabase(countryDbCollection, cca3Code)

	if len(countryObtainedFromDb) != 0 {
		// Checks if the cache is more than ten days old, if not it will return the item from the db.
		if !(time.Since(countryObtainedFromDb["timestamp"].(time.Time)).Hours() > (time.Hour * 24 * 10).Hours()) {
			return countryObtainedFromDb["name"].(string), nil
		}
	}

	url := buildSearchUrl(cca3Code)

	response, errResponse := web_client.GetRequest(url)

	if errResponse != nil {
		return "", errResponse
	}

	country := decodeCountryInfo(response)

	// If the country object is an empty struct.
	if (countryStruct{}) == country {
		return "", nil
	}

	err := database.WriteToDatabase(countryDbCollection, cca3Code, map[string]string{"name": country.Name.Common})

	if err != nil {
		return "", err
	}

	return country.Name.Common, nil
}
