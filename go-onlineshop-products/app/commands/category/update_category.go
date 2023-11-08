package category

import (
	"context"
	domain "onlineshopproduct/domain"
)

type UpdateCategoryHandler struct {
	repo domain.CategoryRepository
}

func NewUpdateCategoryHandler(repo domain.CategoryRepository) UpdateCategoryHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return UpdateCategoryHandler{repo}
}

func (h UpdateCategoryHandler) Handle(ctx context.Context, category domain.Category) error {
	err := h.repo.UpdateCategory(ctx, category)

	if err != nil {
		return err
	}

	return nil
}
