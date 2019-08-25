package controller

import (
	"encoding/json"
	"go-restful-api/database"
	"go-restful-api/models"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := database.Init()

	user := model.User{
		First: "Nefertiti",
		Last:  "Okoh",
	}

	db.Create(&user)

	w.WriteHeader(http.StatusCreated)

	message := Response{
		Message: "Success️",
		Status:  201,
	}

	_ = json.NewEncoder(w).Encode(message)
}
