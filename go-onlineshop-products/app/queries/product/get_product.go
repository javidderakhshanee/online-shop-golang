package product

import (
	"context"
	domain "onlineshopproduct/domain"
)

type GetProductHandler struct {
	repo domain.ProductRepository
}

func NewGetProductHandler(repo domain.ProductRepository) GetProductHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetProductHandler{repo}
}

func (handler GetProductHandler) Handle(ctx context.Context, id int) (domain.Product, error) {
	product, err := handler.repo.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
