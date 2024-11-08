package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0z-cryptik/randomJokesAPI/structs"
)

var HandlerFuncMap = map[string]http.HandlerFunc{
	"/getJoke": jokeHandler,
	"/cron":    cronHandler,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if handler, exists := HandlerFuncMap[path]; exists {
		handler(w, r)
		return
	}

	// sends error response if the endpoint doesn't exist
	errorResponse := structs.ErrorResponse{
		Ok:      false,
		Message: fmt.Sprintf("'%v' this endpoint doesn't exist", path),
	}

	errorResponseJson, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, "An error occurred, please try again", http.StatusInternalServerError)
		fmt.Printf("Error %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(errorResponseJson)
}
