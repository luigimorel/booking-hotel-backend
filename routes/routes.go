package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/controllers"
	"github.com/morelmiles/booking-backend/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes() {
	middleware.InitLogger()

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	// Users
	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/v1/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/api/v1/user/{id}", controllers.UpdateUserById).Methods("PUT")

	// Properties
	router.HandleFunc("/api/v1/properties", controllers.GetProperties).Methods("GET")
	router.HandleFunc("/api/v1/property/{id}", controllers.GetPropertyById).Methods("GET")
	router.HandleFunc("/api/v1/properties", controllers.CreateProperty).Methods("POST")
	router.HandleFunc("/api/v1/property/{id}", controllers.DeletePropertyById).Methods("DELETE")
	router.HandleFunc("/api/v1/property/{id}", controllers.UpdatePropertyById).Methods("PUT")

	http.ListenAndServe(":8080", router)
}
