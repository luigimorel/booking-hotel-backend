package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/morelmiles/booking-backend/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
)

func baseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func Routes() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
	}
	port := os.Getenv("PORT")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", baseRoute)
	router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/v1/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/api/v1/register", controllers.UpdateUserById).Methods("PUT")

	http.ListenAndServe(port, router)
}
