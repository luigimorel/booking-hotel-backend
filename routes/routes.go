package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/controllers"
)

func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	http.ListenAndServe(":8080", router)
}
