package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/morelmiles/booking-backend/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
)

func Routes() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
	}
	port := os.Getenv("PORT")

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/v1/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/api/v1/register", controllers.UpdateUserById).Methods("PUT")

	http.ListenAndServe(":"+port, router)
}
