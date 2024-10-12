package main

import (
	"flashcards/db"
	"flashcards/router"
	"log"
)

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
