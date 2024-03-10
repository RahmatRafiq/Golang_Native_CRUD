package main

import (
	"Golang_CRUD_Native/config"
	"Golang_CRUD_Native/controllers/categorycontroller"
	"Golang_CRUD_Native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/destroy", categorycontroller.Destroy)

	// http.HandleFunc("/categories", categorycontroller.Index)
	// http.HandleFunc("/categories/add", categorycontroller.Add)
	// http.HandleFunc("/categories/edit", categorycontroller.Edit)
	// http.HandleFunc("/categories/destroy", categorycontroller.Destroy)

	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
