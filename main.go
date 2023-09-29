package main

import (
	"log"
	"os"
)

var (
	port     = os.Getenv("PORT")
	temp_dir = "tmp"
)

func main() {
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err.Error())
	}
}
