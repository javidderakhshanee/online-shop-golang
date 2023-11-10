package product

import (
	//"fmt"
	//	"html/template"
	//"log"
	"net/http"
	//"onlineshopresentation/restapi"
)

func StartHandler() {

	http.HandleFunc("/product", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	/*url := fmt.Sprintf(baseXkcdURL, categoryId)

	products, err := restapi.GetRestClient(url)
	if err != nil {
		log.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("/product/product.html"))

	tmpl.Execute(w, products)*/

}
