package category

import (
	"context"
	domain "onlineshopproduct/domain"
)

type DeleteCategoryHandler struct {
	repo domain.CategoryRepository
}

func NewDeleteCategoryHandler(repo domain.CategoryRepository) DeleteCategoryHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return DeleteCategoryHandler{repo}
}

func (h DeleteCategoryHandler) Handle(ctx context.Context, id int) error {
	err := h.repo.DeleteCategory(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
