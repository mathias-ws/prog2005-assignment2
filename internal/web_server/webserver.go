package web_server

import (
	"assignment-2/internal/web_server/handlers"
	"log"
	"net/http"
	"os"
)

// setHandlers sets all the web handlers that the server has.
func setHandlers() {
	http.HandleFunc(statusLocation, handlers.StatusHandler)
	http.HandleFunc(policyLocation, handlers.PolicyHandler)
	http.HandleFunc(casesLocation, handlers.CovidCasesHandler)
	http.HandleFunc(webhookLocation, handlers.NotificationHandler)
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

	log.Println("Webserver started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
