package handlers

import (
	"encoding/json"
	"fmt"
	"go-api-postgress/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Error deleting todo: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if rows > 1 {
		log.Printf("Error deleting: %d", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo deleted with id: %d", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
