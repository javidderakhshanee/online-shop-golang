package category

import (
	"context"
	domain "onlineshopproduct/domain"
)

type AddCategoryHandler struct {
	repo domain.CategoryRepository
}

func NewAddCategoryHandler(repo domain.CategoryRepository) AddCategoryHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return AddCategoryHandler{repo}
}

func (h AddCategoryHandler) Handle(ctx context.Context, category domain.Category) error {
	err := h.repo.AddCategory(ctx, category)

	if err != nil {
		return err
	}

	return nil
}
