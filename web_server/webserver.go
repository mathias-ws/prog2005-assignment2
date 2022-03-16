package web_server

import (
	"assignment-2/status_endpoint"
	"assignment-2/web_server/handlers"
	"log"
	"net/http"
	"os"
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
