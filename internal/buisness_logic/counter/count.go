package counter

import (
	"assignment-2/internal/buisness_logic/country"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
)

// CountUp counts up the number of times the country have been searched.
func CountUp(countryName string) error {
	if len(countryName) == 3 {
		var err error
		countryName, err = country.GetCountryNameFromCca3(countryName)

		if err != nil {
			return err
		}
	}

	database.IncrementCounter(CounterDbCollection, countryName)

	return nil
}

// GetNumberOfTimes gets the number of times a country has been counted.
func GetNumberOfTimes(countryName string) (int, error) {
	if len(countryName) == 3 {
		var err error
		countryName, err = country.GetCountryNameFromCca3(countryName)

		if err != nil {
			return 0, err
		}
	}

	var obtainedCountry structs.CountryCounter
	database.GetDocument(CounterDbCollection, countryName, &obtainedCountry)

	if (structs.CountryCounter{}) == obtainedCountry {
		return 0, custom_errors.GetNoObjectFoundError()
	}

	return obtainedCountry.Count, nil
}
