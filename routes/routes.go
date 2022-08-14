package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
)

func baseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/", baseRoute)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")

	http.ListenAndServe(":8080", router)
}
