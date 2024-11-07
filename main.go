package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/0z-cryptik/randomJokesAPI/handlers"
)

type ResponseStruct struct {
	Ok bool
	Joke string
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handlers.Handler)

	http.ListenAndServe(":" + port, nil)
}