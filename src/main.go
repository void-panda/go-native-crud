package main

import (
	"go-native-crud/config"
	"go-native-crud/src/routes"
	"log"
	"net/http"
)

func main() {
	// connect DB
	config.InitDB()
	log.Println("DB Connected ✅")

	// Routes
	routes.HomeRoutes()
	routes.CategoryRoutes()
	routes.PostRoutes()

	log.Println("Starting server on port 3333")
	http.ListenAndServe(":3333", nil)
}
