package routes

import (
	categorycontroller "go-native-crud/src/controllers/categoryController"
	"net/http"
)

func CategoryRoutes() {
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/store", categorycontroller.Store)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/update", categorycontroller.Update)
	http.HandleFunc("/categories/destroy", categorycontroller.Destroy)
}
