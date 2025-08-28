package postcontroller

import (
	"go-native-crud/src/entities"
	categorymodel "go-native-crud/src/models/categoryModel"
	postmodel "go-native-crud/src/models/postModel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"posts":      postmodel.GetAll(),
		"categories": categorymodel.GetAll(),
	}

	temp, err := template.ParseFiles("src/views/posts/index.html")
	if err != nil {
		panic(err)
	}

	err = temp.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	categoryId, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		panic(err)
	}
	post := entities.Post{
		Title: r.FormValue("title"),
		Category: entities.Category{
			Id: uint(categoryId),
		},
		Content: r.FormValue("content"),
	}

	ok := postmodel.Create(post)
	if !ok {
		temp, err := template.ParseFiles("src/views/posts/index.html")
		if err != nil {
			panic(err)
		}

		err = temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	postId, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err)
	}
	data := map[string]any{
		"post":       postmodel.GetById(postId),
		"categories": categorymodel.GetAll(),
	}

	temp, err := template.ParseFiles("src/views/posts/edit.html")
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
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	postId, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err)
	}

	categoryId, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		panic(err)
	}

	post := entities.Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
		Category: entities.Category{
			Id: uint(categoryId),
		},
		Updated_at: time.Now(),
	}

	if ok := postmodel.Update(postId, post); !ok {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/posts", http.StatusSeeOther)

}

func Destroy(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	err = postmodel.Delete(postId)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}
