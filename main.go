package main

import (
	"assignment-2/database"
	"assignment-2/web_server"
)

// main starts the program.
func main() {
	database.InitDB()
	web_server.StartWebServer()
}
