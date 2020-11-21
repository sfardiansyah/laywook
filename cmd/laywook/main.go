package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sfardiansyah/laywook/pkg/auth"
	"github.com/sfardiansyah/laywook/pkg/http/rest"
	"github.com/sfardiansyah/laywook/pkg/storage/mongodb"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := mongodb.NewStorage()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	a := auth.Service(s)
	r := rest.Handler(a)

	fmt.Printf("Starting beanie at http://localhost:%s/\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
