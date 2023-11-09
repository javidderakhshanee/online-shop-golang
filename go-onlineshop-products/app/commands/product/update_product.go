package product

import (
	"context"
	domain "onlineshopproduct/domain"
)

type UpdateProductHandler struct {
	repo domain.ProductRepository
}

func NewUpdateProductHandler(repo domain.ProductRepository) UpdateProductHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return UpdateProductHandler{repo}
}

func (h UpdateProductHandler) Handle(ctx context.Context, product domain.Product) error {
	err := h.repo.UpdateProduct(ctx, product)

	if err != nil {
		return err
	}

	return nil
}
