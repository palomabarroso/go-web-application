package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/palomabarroso/go-web-application/routes"
)

func main() {
	godotenv.Load()
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
