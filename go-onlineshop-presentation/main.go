package main

import (
	"log"
	"net/http"
	helper "onlineshopresentation/global"
	"onlineshopresentation/product"
	"os"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", indexHandler)

	product.StartHandler()

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "localhost:8086"
	}

	log.Fatal(http.ListenAndServe(address, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &helper.MasterPageInformation[string]{
		Title: "Online Shop | Martket Place",
		Body:  "Welcome to the Online shop.",
	}

	helper.RunTemplate("otherpages/main.html", "main", data, w)
}
