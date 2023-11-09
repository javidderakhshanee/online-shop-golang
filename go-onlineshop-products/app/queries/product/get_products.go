package product

import (
	"context"
	domain "onlineshopproduct/domain"
)

type GetProductsHandler struct {
	repo domain.ProductRepository
}

func NewGetProductsHandler(repo domain.ProductRepository) GetProductsHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetProductsHandler{repo}
}

func (handler GetProductsHandler) Handle(ctx context.Context, categoryId int) ([]domain.Product, error) {
	products, err := handler.repo.GetProducts(ctx, categoryId)

	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}
