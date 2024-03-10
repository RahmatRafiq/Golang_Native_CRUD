package categorycontroller

import (
	"Golang_CRUD_Native/models/categorymodel"
	"html/template"
	"net/http"
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

}

func Edit(W http.ResponseWriter, r *http.Request) {

}

func Destroy(W http.ResponseWriter, r *http.Request) {

}
