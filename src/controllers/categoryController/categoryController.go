package categorycontroller

import (
	"go-native-crud/src/entities"
	categorymodel "go-native-crud/src/models/categoryModel"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Index
func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"categories": categorymodel.GetAll(),
	}
	temp, err := template.ParseFiles("src/views/categories/index.html")
	if err != nil {
		panic(err.Error())
	}

	err = temp.Execute(w, data)
	if err != nil {
		panic(err.Error())
	}
}

// Create
func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	category := entities.Category{
		Name:       r.FormValue("name"),
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	if ok := categorymodel.Create(category); !ok {
		temp, _ := template.ParseFiles("src/views/categories/index.html")
		err := temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	categoryId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	data := map[string]any{
		"category": categorymodel.GetById(categoryId),
	}

	temp, err := template.ParseFiles("src/views/categories/edit.html")
	if err != nil {
		panic(err)
	}

	err = temp.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("=========> Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	categoryId, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err.Error())
	}

	category := entities.Category{
		Name:       r.FormValue("name"),
		Updated_at: time.Now(),
	}

	if ok := categorymodel.Update(categoryId, category); !ok {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)

}

func Destroy(w http.ResponseWriter, r *http.Request) {
	categoryId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	if err := categorymodel.Delete(categoryId); err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
