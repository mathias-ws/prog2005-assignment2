package json_parser

import (
	"assignment-2/model"
	"encoding/json"
	"log"
	"net/http"
)

// DecodePolicyInfo decodes the policy info into the PolicyInputFromApi struct.
func DecodePolicyInfo(httpResponse *http.Response) model.PolicyInputFromApi {
	decoder := json.NewDecoder(httpResponse.Body)
	var list model.PolicyInputFromApi

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&list); err != nil {
		log.Println(err)
	}

	return list
}
