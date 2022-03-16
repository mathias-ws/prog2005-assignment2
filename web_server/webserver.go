package web_server

import (
	"log"
	"net/http"
	"os"
	"status_endpoint"
	"web_server/v1.0.0/handlers"
)

// setHandlers sets all the web handlers that the server has.
func setHandlers() {
	http.HandleFunc(statusLocation, handlers.StatusHandler)
	http.HandleFunc(policyLocation, handlers.PolicyHandler)
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
