package main

import (
	"fmt"
	"net/http"
	"github.com/0z-cryptik/randomJokesAPI/jokes"
	"encoding/json"
)

type ResponseStruct struct {
	Ok bool
	Joke string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		response := ResponseStruct{
			Ok: true,
			Joke: jokes.Jokes[0],
		}

		responseJson, err := json.Marshal(response)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			fmt.Printf("some error %v\n", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(responseJson)
	})

	http.ListenAndServe(":8080", nil)
}