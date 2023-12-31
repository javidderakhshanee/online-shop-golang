package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"onlineshopbasket/app"
	controller "onlineshopbasket/interfaces"
	"os"
)

func main() {

	r := mux.NewRouter()
	ctx := context.Background()
	app := app.NewApplication(ctx)
	c := controller.NewBasketController(app)

	r.HandleFunc("/baskets/{id}", c.GetBasket).Methods("GET")
	r.HandleFunc("/baskets/{id}", c.UpdateBasket).Methods("PUT")
	r.HandleFunc("/baskets/{id}", c.DeleteBasket).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	fmt.Println("Start listening on port", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

}
