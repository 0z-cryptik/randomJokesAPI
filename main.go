package main

import (
	"net/http"
	"github.com/0z-cryptik/randomJokesAPI/handlers"
	"math/rand"
	"time"
)

type ResponseStruct struct {
	Ok bool
	Joke string
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handlers.Handler)

	http.ListenAndServe(":8080", nil)
}