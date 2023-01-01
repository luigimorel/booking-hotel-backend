package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/config"
	"github.com/morelmiles/booking-backend/models"
)

func checkIfPropertyExists(propertyId string) bool {
	var property models.Property
	config.DB.First(&property, propertyId)

	return property.ID != 0
}

// GetProperties - Fetches a list of all properties.
func GetProperties(w http.ResponseWriter, r *http.Request) {

	var properties []models.Property

	config.DB.Find(&properties)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&properties)
}

// GetPropertyById - Returns a single property with the value of the ID specified.
func GetPropertyById(w http.ResponseWriter, r *http.Request) {
	propertyId := mux.Vars(r)["id"]
	if !checkIfPropertyExists(propertyId) {
		json.NewEncoder(w).Encode("property not found!")
		return
	}
	var property models.Property
	config.DB.First(&property, propertyId)
	json.NewEncoder(w).Encode(property)
}

// CreateProperty - Creates a new property
func CreateProperty(w http.ResponseWriter, r *http.Request) {

	var job models.Property
	var err error

	json.NewDecoder(r.Body).Decode(&job)

	config.DB.Create(&job)

	if err != nil {
		log.Panic(err)
	} else {
		json.NewEncoder(w).Encode(&job)
	}
}

// UpdatePropertyById -  Updates a single property by the ID specified
func UpdatePropertyById(w http.ResponseWriter, r *http.Request) {
	propertyId := mux.Vars(r)["id"]
	if !checkIfPropertyExists(propertyId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("property not found!")
		return
	}

	var property models.Property

	config.DB.First(&property, propertyId)
	json.NewDecoder(r.Body).Decode(&property)
	config.DB.Save(&property)
	json.NewEncoder(w).Encode(property)
}

// DeletePropertyById - Updates a single property by the ID specified.
func DeletePropertyById(w http.ResponseWriter, r *http.Request) {
	propertyId := mux.Vars(r)["id"]
	if !checkIfPropertyExists(propertyId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("property not found!")
		return
	}

	var property models.Property
	config.DB.Delete(&property, propertyId)
	json.NewEncoder(w).Encode(property)
}
