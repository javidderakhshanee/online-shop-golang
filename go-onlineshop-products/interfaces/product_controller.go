package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"onlineshopproduct/app"
	domain "onlineshopproduct/domain"
	"strconv"
)

type ProductController struct {
	app app.Application
}

func NewProductController(app app.Application) ProductController {
	return ProductController{app}
}

func (controller ProductController) GetProducts(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	categoryId, err := strconv.Atoi(params["categoryId"])
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if categoryId == 0 {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	products, err := controller.app.Queries.GetProducts.Handle(request.Context(), categoryId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productsJson, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(productsJson))
}

func (controller ProductController) GetProduct(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if id == 0 {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	Product, err := controller.app.Queries.GetProduct.Handle(request.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ProductJson, _ := json.Marshal(Product)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(ProductJson))
}

func (controller ProductController) AddProduct(w http.ResponseWriter, request *http.Request) {

	var newProduct domain.Product

	err := json.NewDecoder(request.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	controller.app.Commands.AddProduct.Handle(request.Context(), newProduct)

	w.WriteHeader(http.StatusNoContent)
}

func (controller ProductController) UpdateProduct(w http.ResponseWriter, request *http.Request) {

	var Product domain.Product

	err := json.NewDecoder(request.Body).Decode(&Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	controller.app.Commands.UpdateProduct.Handle(request.Context(), Product)

	w.WriteHeader(http.StatusNoContent)
}

func (controller ProductController) DeleteProduct(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if id == 0 {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controller.app.Commands.DeleteProduct.Handle(request.Context(), id)

	w.WriteHeader(http.StatusNoContent)
}
