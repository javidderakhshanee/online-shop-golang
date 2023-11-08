package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"onlineshopproduct/app"
	controller "onlineshopproduct/interfaces"

	"github.com/gorilla/mux"
)

func addCategoryController(router *mux.Router, application app.Application) {
	controller := controller.NewCategoryController(application)

	router.HandleFunc("/products/categories", controller.GetCategories).Methods("GET")
	router.HandleFunc("/products/category/{id}", controller.GetCategory).Methods("GET")
	router.HandleFunc("/products/category", controller.AddCategory).Methods("POST")
	router.HandleFunc("/products/category", controller.UpdateCategory).Methods("PUT")
	router.HandleFunc("/products/category/{id}", controller.DeleteCategory).Methods("DELETE")
}

func main() {

	router := mux.NewRouter()
	ctx := context.Background()
	app := app.NewApplication(ctx)

	addCategoryController(router, app)

	port := ":8001"

	fmt.Println("Start listening on port", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}

}
