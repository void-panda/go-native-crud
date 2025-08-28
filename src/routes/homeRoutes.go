package routes

import (
	homecontroller "go-native-crud/src/controllers/homeController"
	"net/http"
)

func HomeRoutes() {
	http.HandleFunc("/", homecontroller.Index)
}
