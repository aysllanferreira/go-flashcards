package main

import (
	"flashcards/db"
	"flashcards/router"
	"log"
)

// @title Flashcards API
// @version 1.0
// @description This is my flashcards API.
// @host localhost:8080
// @BasePath /
func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	r := router.SetupRouter(dbConnection)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
