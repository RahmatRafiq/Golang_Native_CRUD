package categorycontroller

import (
	"Golang_CRUD_Native/entities"
	"Golang_CRUD_Native/models/categorymodel"
	"html/template"
	"net/http"
	"time"
)

func Index(W http.ResponseWriter, r *http.Request) {
	categories := categorymodel.Getall()
	data := map[string]interface{}{
		"categories": categories,
	}
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(W, data)

}

func Add(W http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(W, nil)

	}

	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(W, nil)

			http.Redirect(W, r, "/categories", http.StatusSeeOther)
		}
	}

}

func Edit(W http.ResponseWriter, r *http.Request) {

}

func Destroy(W http.ResponseWriter, r *http.Request) {

}
