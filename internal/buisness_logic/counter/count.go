package counter

import (
	"assignment-2/internal/buisness_logic/country"
	"assignment-2/internal/database"
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
