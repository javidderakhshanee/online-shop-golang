package queries

import (
	"context"
	domain "onlineshopbasket/domain"
)

type GetBasketHandler struct {
	repo domain.Repository
}

func NewGetBasketHandler(repo domain.Repository) GetBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetBasketHandler{repo}
}

func (h GetBasketHandler) Handle(ctx context.Context, id string) (domain.BasketHeader, error) {
	b, err := h.repo.GetBasket(ctx, id)
	if err != nil {
		return domain.BasketHeader{}, err
	}

	return b, nil
}
