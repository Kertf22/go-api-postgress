package handlers

import (
	"encoding/json"
	"go-api-postgress/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error getting todos: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
