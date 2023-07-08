package handlers

import (
	"action-logger/pkg/models"
	"encoding/json"
	"net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var submitLogger models.SubmitLogger
	err := json.NewDecoder(r.Body).Decode(&submitLogger)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Action logger submitted successfully"))
}
