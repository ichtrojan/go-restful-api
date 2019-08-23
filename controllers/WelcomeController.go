package controller

import (
	"encoding/json"
	"net/http"
)

type Welcome struct {
	Message string
	Status  int64
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	message := Welcome{
		Message: "Welcome ❤️",
		Status:  200,
	}

	_ = json.NewEncoder(w).Encode(message)
}
