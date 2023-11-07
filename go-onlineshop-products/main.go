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

	router.HandleFunc("/products/category", controller.AddCategory).Methods("POST")
}

func main() {

	router := mux.NewRouter()
	ctx := context.Background()
	app := app.NewApplication(ctx)

	addCategoryController(router, app)

	port := ":9002"

	fmt.Println("Start listening on port", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}

}
