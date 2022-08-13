package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/config"
	"github.com/morelmiles/booking-backend/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	config.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User

	config.DB.First(&user, params["id"])

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	newUser := config.DB.Create(&user)
	err := newUser.Error

	if err != nil {
		log.Panic(err)
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}
