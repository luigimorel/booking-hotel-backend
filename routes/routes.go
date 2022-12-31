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

	// Swagger docs
	router.PathPrefix("/api/v1/swagger").Handler(httpSwagger.WrapHandler)

	// Home
	router.HandleFunc("/", controllers.Home).Methods("GET")

	// Auth
	router.HandleFunc("/api/v1/auth/login", middleware.SetMiddlewareJSON(controllers.Login)).Methods("POST")

	// Users
	router.HandleFunc("/api/v1/users", middleware.SetMiddlewareJSON(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", middleware.SetMiddlewareJSON(controllers.GetUserById)).Methods("GET")
	router.HandleFunc("/api/v1/register", middleware.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", middleware.SetMiddlewareAuthentication(controllers.DeleteUserById)).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", middleware.SetMiddlewareAuthentication(controllers.UpdateUserById)).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}/properties", middleware.SetMiddlewareJSON(controllers.GetAllPropertiesByUser)).Methods("GET")

	// Properties
	router.HandleFunc("/api/v1/properties", middleware.SetMiddlewareJSON(controllers.GetProperties)).Methods("GET")
	router.HandleFunc("/api/v1/properties/{id}", middleware.SetMiddlewareAuthentication(controllers.GetPropertyById)).Methods("GET")
	router.HandleFunc("/api/v1/properties", middleware.SetMiddlewareAuthentication(controllers.CreateProperty)).Methods("POST")
	router.HandleFunc("/api/v1/properties/{id}", middleware.SetMiddlewareAuthentication(controllers.DeletePropertyById)).Methods("DELETE")
	router.HandleFunc("/api/v1/properties/{id}", middleware.SetMiddlewareAuthentication(controllers.UpdatePropertyById)).Methods("PUT")

	// Server port
	http.ListenAndServe(":8080", router)
}
