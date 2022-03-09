package web_server

import (
	"assignment-2/constants"
	"log"
	"net/http"
	"os"
)

// setHandlers sets all the web handlers that the server has.
func setHandlers() {
	//TODO add handlers
}

// StartWebServer starts the webserver for the api.
func StartWebServer() {
	port := os.Getenv("PORT")

	// Checks if the port env variable is set, if not it will set the default port.
	if port == "" {
		log.Println("Port variable not set, using default: " + constants.PORT)
		port = constants.PORT
	}

	setHandlers()

	log.Println("Webserver started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
