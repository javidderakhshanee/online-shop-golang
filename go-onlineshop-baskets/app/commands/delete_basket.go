package commands

import (
	"context"
	domain "onlineshopbasket/domain"
)

type DeleteBasketHandler struct {
	repo domain.Repository
}

func NewDeleteBasketHandler(repo domain.Repository) DeleteBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return DeleteBasketHandler{repo}
}

func (h DeleteBasketHandler) Handle(ctx context.Context, id string, b domain.BasketHeader) error {
	err := h.repo.DeleteBasket(ctx, id, b)

	if err != nil {
		return err
	}

	return nil
}
