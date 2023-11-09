package product

import (
	"context"
	domain "onlineshopproduct/domain"
)

type DeleteProductHandler struct {
	repo domain.ProductRepository
}

func NewDeleteProductHandler(repo domain.ProductRepository) DeleteProductHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return DeleteProductHandler{repo}
}

func (h DeleteProductHandler) Handle(ctx context.Context, id int) error {
	err := h.repo.DeleteProduct(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
