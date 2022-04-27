package country

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/json_parsing"
	"assignment-2/internal/structs"
	"assignment-2/internal/web_client"
	"log"
	"strings"
	"time"
)

// buildSearchUrl builds a search url for the country api.
func buildSearchUrl(cca3Code string) string {
	urlToSearch := strings.Builder{}
	urlToSearch.WriteString(constants.CountryApiUrl)
	urlToSearch.WriteString(constants.CountryApiPage)
	urlToSearch.WriteString(countryApiCodeParam)
	urlToSearch.WriteString(cca3Code)
	urlToSearch.WriteString(countryApiFieldParam)

	return urlToSearch.String()
}

// GetCountryNameFromCca3 converts the inputted cca3 code into a full country name.
func GetCountryNameFromCca3(cca3Code string) (string, error) {
	var countryObtainedFromDb structs.CountryInfo
	database.GetDocument(constants.CountryDbCollection, cca3Code, &countryObtainedFromDb)

	if (structs.CountryInfo{}) != countryObtainedFromDb {
		// Checks if the cache is more than ten days old, if not it will return the item from the db.
		if !(time.Since(countryObtainedFromDb.TimeStamp).Hours() > (time.Hour * 24 * 50).Hours()) {
			return countryObtainedFromDb.Common, nil
		}
	}

	url := buildSearchUrl(cca3Code)

	response, errResponse := web_client.GetRequest(url)

	if errResponse != nil {
		return "", errResponse
	}

	var country []structs.CountryNameStruct
	json_parsing.Decode(response, &country)

	if len(country) == 0 {
		return "", custom_errors.GetFailedToDecode()
	}

	go func() {
		err := database.WriteDocument(constants.CountryDbCollection, cca3Code, map[string]string{"name": country[0].Name.Common})

		if err != nil {
			log.Printf("Error writing to db: %v", err)
		}
	}()

	return country[0].Name.Common, nil
}
