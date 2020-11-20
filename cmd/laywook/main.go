package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Hello World!")

	r := http.NewServeMux()

	fmt.Printf("Starting beanie at http://localhost:%s/\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
