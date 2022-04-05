package stubs

import (
	"assignment-2/internal/structs"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// CountryHandler returns stub data for the country endpoint
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	timestamp, _ := time.Parse("2006-01-02", "2022-03-20")

	jsonDataNor := []structs.CountryNameStruct{{
		Name: structs.CountryInfo{
			Common:    "Norway",
			TimeStamp: timestamp,
		},
	},
	}

	jsonDataNld := []structs.CountryNameStruct{{
		Name: structs.CountryInfo{
			Common:    "Netherlands",
			TimeStamp: timestamp,
		},
	},
	}

	jsonDataLva := []structs.CountryNameStruct{{
		Name: structs.CountryInfo{
			Common:    "Latvia",
			TimeStamp: timestamp,
		},
	},
	}

	jsonDataFra := []structs.CountryNameStruct{{
		Name: structs.CountryInfo{
			Common:    "France",
			TimeStamp: timestamp,
		},
	},
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)

	if strings.Contains(r.URL.RawQuery, "nor") {
		err := encoder.Encode(jsonDataNor)
		if err != nil {
			return
		}
		return
	} else if strings.Contains(r.URL.RawQuery, "nld") {
		err := encoder.Encode(jsonDataNld)
		if err != nil {
			return
		}
		return
	} else if strings.Contains(r.URL.RawQuery, "lva") {
		err := encoder.Encode(jsonDataLva)
		if err != nil {
			return
		}
		return
	} else if strings.Contains(r.URL.RawQuery, "fra") {
		err := encoder.Encode(jsonDataFra)
		if err != nil {
			return
		}
		return
	}

}
