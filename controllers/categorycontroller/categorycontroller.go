package categorycontroller

import (
	"Golang_CRUD_Native/entities"
	"Golang_CRUD_Native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
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

// func Edit(W http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		temp, err := template.ParseFiles("views/category/edit.html")
// 		if err != nil {
// 			panic(err)
// 		}

// 		idString := r.URL.Query().Get("id")
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			panic(err)
// 		}

// 		category := categorymodel.Detail(id)
// 		data := map[string]interface{}{
// 			"category": category,
// 		}
// 		temp.Execute(W, data)
// 	}
// }

func Edit(W http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}
		temp.Execute(W, data)
	}

	if r.Method == "POST" {
		var category entities.Category
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Redirect(W, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(W, r, "/categories", http.StatusSeeOther)
	}
}

func Destroy(W http.ResponseWriter, r *http.Request) {

}
