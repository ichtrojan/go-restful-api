package controller

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	message := Response{
		Message: "Welcome ❤️",
		Status:  200,
	}

	_ = json.NewEncoder(w).Encode(message)
}
