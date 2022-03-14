package web_server

import (
	"assignment-2/constants"
	"assignment-2/logic"
	"assignment-2/web_server/handlers"
	"log"
	"net/http"
	"os"
)

// setHandlers sets all the web handlers that the server has.
func setHandlers() {
	http.HandleFunc(constants.STATUS_LOCATION, handlers.StatusHandler)
	http.HandleFunc(constants.POLICY_LOCATION, handlers.PolicyHandler)
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
	logic.SetStartTime()

	log.Println("Webserver started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
