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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

// GetUserById ... Get one user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		json.NewEncoder(w).Encode("user not found!")
		return
	}
	var user models.User
	config.DB.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func checkIfUserExists(userId string) bool {
	var user models.User
	config.DB.First(&user, userId)

	return user.ID != 0
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found!")
		return
	}
	var user models.User
	config.DB.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	var user models.User
	config.DB.Delete(&user, userId)
	json.NewEncoder(w).Encode(user)
}
