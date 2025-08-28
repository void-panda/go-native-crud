package homecontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("src/views/home/index.html")
	if err != nil {
		panic(err.Error())
	}

	err = temp.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
