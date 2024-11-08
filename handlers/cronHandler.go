package handlers
import "net/http"

func cronHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}