package main

import (
	"html/template"
	"log"
	"net/http"
	//"onlineshopresentation/product"
	//"strings"
)

func main() {

	http.HandleFunc("/", indexHandler)

	//product.StartHandler()

	log.Fatal(http.ListenAndServe("localhost:8086", nil))
}

var templates = template.Must(template.ParseFiles("templates/header.tmpl", "templates/footer.tmpl", "otherpages/main.html"))

type Index struct {
	Title string
	Body  string
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Online Shop | Martket Place",
		Body:  "Welcome to the Online shop.",
	}

	display(w, "main", data)
}
