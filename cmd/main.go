package main

import (
	"assignment-2/internal/database"
	"assignment-2/internal/web_server"
)

// main starts the program.
func main() {
	database.InitDB("auth.json")
	defer database.CloseFirestore()
	web_server.StartWebServer()
}
