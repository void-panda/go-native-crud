package routes

import (
	postcontroller "go-native-crud/src/controllers/postController"
	"net/http"
)

func PostRoutes() {
	http.HandleFunc("/posts", postcontroller.Index)
	http.HandleFunc("/posts/store", postcontroller.Store)
	http.HandleFunc("/posts/edit", postcontroller.Edit)
	http.HandleFunc("/posts/update", postcontroller.Update)
	http.HandleFunc("/posts/destroy", postcontroller.Destroy)
}
