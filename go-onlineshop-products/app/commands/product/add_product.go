package product

import (
	"context"
	domain "onlineshopproduct/domain"
)

type AddProductHandler struct {
	repo domain.ProductRepository
}

func NewAddProductHandler(repo domain.ProductRepository) AddProductHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return AddProductHandler{repo}
}

func (h AddProductHandler) Handle(ctx context.Context, product domain.Product) error {
	err := h.repo.AddProduct(ctx, product)

	if err != nil {
		return err
	}

	return nil
}
