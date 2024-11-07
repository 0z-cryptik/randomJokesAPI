package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0z-cryptik/randomJokesAPI/structs"
)

func Handler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/getJoke" {
		errorResponse := structs.ErrorResponse{
			Ok: false,
			Message: fmt.Sprintf("'%v' this endpoint doesn't exist", r.URL.Path),
		}

		errorResponseJson, err := json.Marshal(errorResponse)

		if err != nil {
			http.Error(w, "An error occured, please try again", http.StatusInternalServerError)
			fmt.Printf("Error %v\n", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorResponseJson)
		return
	}

	jokeHandler(w, r)
}