package interfaces

import (
	"encoding/json"
	"net/http"
	"onlineshopbasket/app"
	domain "onlineshopbasket/domain"

	"github.com/gorilla/mux"
)

type BasketController struct {
	app app.Application
}

func NewBasketController(app app.Application) BasketController {
	return BasketController{app}
}

func (c BasketController) GetBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	b, err := c.app.Queries.GetBasket.Handle(r.Context(), params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basketJson, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(basketJson))
}

func (c BasketController) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var b domain.BasketHeader

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.app.Commands.UpdateBasket.Handle(r.Context(), params["id"], b)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(params["id"]))
}

func (c BasketController) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	err := c.app.Commands.DeleteBasket.Handle(r.Context(), params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
