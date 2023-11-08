package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"onlineshopproduct/app"
	domain "onlineshopproduct/domain"
	"strconv"
)

type CategoryController struct {
	app app.Application
}

func NewCategoryController(app app.Application) CategoryController {
	return CategoryController{app}
}

func (controller CategoryController) GetCategories(w http.ResponseWriter, request *http.Request) {

	var categories []domain.Category

	categories, err := controller.app.Queries.GetCategories.Handle(request.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoriesJson, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(categoriesJson))
}

func (controller CategoryController) GetCategory(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if id == 0 {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	category, err := controller.app.Queries.GetCategory.Handle(request.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryJson, _ := json.Marshal(category)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(categoryJson))
}

func (controller CategoryController) AddCategory(w http.ResponseWriter, request *http.Request) {

	var newCategory domain.Category

	err := json.NewDecoder(request.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	controller.app.Commands.AddCategory.Handle(request.Context(), newCategory)

	w.WriteHeader(http.StatusNoContent)
}

func (controller CategoryController) UpdateCategory(w http.ResponseWriter, request *http.Request) {

	var category domain.Category

	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	controller.app.Commands.UpdateCategory.Handle(request.Context(), category)

	w.WriteHeader(http.StatusNoContent)
}

func (controller CategoryController) DeleteCategory(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if id == 0 {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controller.app.Commands.DeleteCategory.Handle(request.Context(), id)

	w.WriteHeader(http.StatusNoContent)
}
