package commands

import (
	"context"
	domain "onlineshopbasket/domain"
)

type UpdateBasketHandler struct {
	repo domain.Repository
}

func NewUpdateBasketHandler(repo domain.Repository) UpdateBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return UpdateBasketHandler{repo}
}

func (h UpdateBasketHandler) Handle(ctx context.Context, id string, b domain.BasketHeader) error {
	err := h.repo.UpdateBasket(ctx, id, b)

	if err != nil {
		return err
	}

	return nil
}
