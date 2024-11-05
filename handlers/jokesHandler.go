package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0z-cryptik/randomJokesAPI/functions"
	"github.com/0z-cryptik/randomJokesAPI/structs"
)

func jokeHandler(w http.ResponseWriter, r *http.Request){
	response := structs.ResponseStruct{
		Ok: true,
		Joke: functions.GetRandomJoke(),
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "An error occured, please try again", http.StatusInternalServerError)
		fmt.Printf("Error %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(responseJson)
}