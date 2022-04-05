package stubs

import (
	"assignment-2/internal/structs"
	"encoding/json"
	"net/http"
	"time"
)

// CountryHandler returns stub data for the country endpoint
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	timestamp, _ := time.Parse("2006-01-02", "2022-03-20")

	jsonData := []structs.CountryNameStruct{{
		Name: structs.CountryInfo{
			Common:    "Norway",
			TimeStamp: timestamp,
		},
	},
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(jsonData)
	if err != nil {
		return
	}
}
