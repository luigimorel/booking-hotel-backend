package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/config"
	"github.com/morelmiles/booking-backend/models"
)

func GetHouseRules(w http.ResponseWriter, r *http.Request) {

	var houseRules []models.HouseRule

	config.DB.Find(&houseRules)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&houseRules)
}

func GetHouseRuleById(w http.ResponseWriter, r *http.Request) {
	ruleId := mux.Vars(r)["id"]

	var houseRule models.User
	config.DB.First(&houseRule, ruleId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houseRule)
}

func checkIfHouseRuleExists(houseRuleId string) bool {
	var houseRule models.HouseRule
	config.DB.First(&houseRule, houseRuleId)

	return houseRule.ID != 0
}

func CreateHouseRule(w http.ResponseWriter, r *http.Request) {
	var houseRule models.HouseRule
	var err error

	json.NewDecoder(r.Body).Decode(&houseRule)

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&houseRule)
}

func UpdateHouseRuleById(w http.ResponseWriter, r *http.Request) {
	houseRuleId := mux.Vars(r)["id"]
	if !checkIfHouseRuleExists(houseRuleId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("house rule not found!")
		return
	}

	var houseRule models.HouseRule

	config.DB.First(&houseRule, houseRuleId)
	json.NewDecoder(r.Body).Decode(&houseRule)
	config.DB.Save(&houseRule)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houseRule)
}

func DeleteHouseRuleById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	houseRuleId := mux.Vars(r)["id"]
	if !checkIfHouseRuleExists(houseRuleId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("house rule not found!")
		return
	}

	var houseRule models.HouseRule
	config.DB.Delete(&houseRule, houseRuleId)
	json.NewEncoder(w).Encode(houseRule)
}
