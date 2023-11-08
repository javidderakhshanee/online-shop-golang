package category

import (
	"context"
	domain "onlineshopproduct/domain"
)

type GetCategoryHandler struct {
	repo domain.CategoryRepository
}

func NewGetCategoryHandler(repo domain.CategoryRepository) GetCategoryHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetCategoryHandler{repo}
}

func (handler GetCategoryHandler) Handle(ctx context.Context, id int) (domain.Category, error) {
	category, err := handler.repo.GetCategory(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}
