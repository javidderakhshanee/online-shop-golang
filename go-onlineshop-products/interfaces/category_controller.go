package interfaces

import (
	"encoding/json"
	"net/http"
	"onlineshopproduct/app"
	domain "onlineshopproduct/domain"
)

type CategoryController struct {
	app app.Application
}

func NewCategoryController(app app.Application) CategoryController {
	return CategoryController{app}
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
