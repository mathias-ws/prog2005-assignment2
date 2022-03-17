package web_client

import (
	"assignment-2/custom_errors"
	"log"
	"net/http"
)

// getResponse method that takes a request and gets a response from the webpage.
func getResponse(request *http.Request) (*http.Response, error) {
	// Instantiate the webClient
	webClient := &http.Client{}

	// Setting timeout for web web_client
	webClient.Timeout = clientTimeout

	// Sending the request
	response, errorFromResponse := webClient.Do(request)

	if errorFromResponse != nil {
		log.Println("Error in the response:", errorFromResponse.Error())
		return nil, custom_errors.GetUnableToReachBackendApisError()
	}

	return response, nil
}

// createRequest creates a request with the wanted method and url.
func createRequest(url string, method string) (*http.Request, error) {
	request, errorFromRequest := http.NewRequest(method, url, nil)

	if errorFromRequest != nil {
		log.Println("Error when creating the request:", errorFromRequest.Error())
		return nil, custom_errors.GetUnableToReachBackendApisError()
	}

	return request, nil
}

// GetRequest sends a get request to a given webpage.
func GetRequest(url string) (*http.Response, error) {
	request, err := createRequest(url, http.MethodGet)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")

	return getResponse(request)
}

// GetStatusCode gets status code from a web page.
func GetStatusCode(url string) (int, error) {
	request, err := createRequest(url, http.MethodHead)

	if err != nil {
		return 0, err
	}

	response, err := getResponse(request)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}
