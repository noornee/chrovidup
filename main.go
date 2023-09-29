package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err.Error())
	}
}
