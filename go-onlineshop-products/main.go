package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	app "onlineshopproduct/app"
	healthchecker "onlineshopproduct/healthchecker"
	controller "onlineshopproduct/interfaces"
	"os"

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

func addProductController(router *mux.Router, application app.Application) {

	controller := controller.NewProductController(application)

	router.HandleFunc("/products/{categoryId}", controller.GetProducts).Methods("GET")
	router.HandleFunc("/products/product/{id}", controller.GetProduct).Methods("GET")
	router.HandleFunc("/products/product", controller.AddProduct).Methods("POST")
	router.HandleFunc("/products/product", controller.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/product/{id}", controller.DeleteProduct).Methods("DELETE")
}

func main() {

	router := mux.NewRouter()
	ctx := context.Background()
	app := app.NewApplication(ctx)

	healthchecker.AddHealthChecker()

	addCategoryController(router, app)

	addProductController(router, app)

	port := os.Getenv("API_PORT")
	fmt.Println("Environment port", port)

	if port == "" {
		port = "8005"
	}

	fmt.Println("Start listening on port", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}
