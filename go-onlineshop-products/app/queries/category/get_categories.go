package category

import (
	"context"
	domain "onlineshopproduct/domain"
)

type GetCategoriesHandler struct {
	repo domain.CategoryRepository
}

func NewGetCategoriesHandler(repo domain.CategoryRepository) GetCategoriesHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetCategoriesHandler{repo}
}

func (handler GetCategoriesHandler) Handle(ctx context.Context) ([]domain.Category, error) {
	categories, err := handler.repo.GetCategories(ctx)
	if err != nil {
		return []domain.Category{}, err
	}

	return categories, nil
}
