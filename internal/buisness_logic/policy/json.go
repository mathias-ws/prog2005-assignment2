package policy

import (
	"encoding/json"
	"log"
	"net/http"
)

// decodePolicyInfo decodes the policy info into the PolicyInputFromApi struct.
func decodePolicyInfo(httpResponse *http.Response) policyInputFromApi {
	decoder := json.NewDecoder(httpResponse.Body)
	var policyInfo policyInputFromApi

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&policyInfo); err != nil {
		log.Println(err)
	}

	return policyInfo
}
