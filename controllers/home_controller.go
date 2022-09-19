package controllers

import (
	"fmt"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API. Time is %s", time.Now())
}
