package web_server

import (
	"assignment-2/custom_errors"
	"assignment-2/status_endpoint"
	"assignment-2/web_server/json_parsing"
	"log"
	"net/http"
	"os"
)

// StatusHandler checks for the http method and handles the error appropriately.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestStatus(w)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

func handleGetRequestStatus(w http.ResponseWriter) {
	err := json_parsing.Encode(w, status_endpoint.GetStatusInfo())

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}

// setHandlers sets all the web handlers that the server has.
func setHandlers() {
	http.HandleFunc(statusLocation, StatusHandler)
	//http.HandleFunc(policyLocation, handlers.PolicyHandler)
}

// StartWebServer starts the webserver for the api.
func StartWebServer() {
	port := os.Getenv("PORT")

	// Checks if the port env variable is set, if not it will set the default port.
	if port == "" {
		log.Println("Port variable not set, using default: " + defaultPort)
		port = defaultPort
	}

	setHandlers()
	status_endpoint.SetStartTime()

	log.Println("Webserver started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
