package main

import (
	"github.com/morelmiles/booking-backend/config"
	"github.com/morelmiles/booking-backend/routes"
)

func main() {
	config.Config()
	routes.Routes()
}
