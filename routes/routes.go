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

	router := mux.NewRouter()
	router.HandleFunc("/", baseRoute)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")

	http.ListenAndServe(port, router)
}
