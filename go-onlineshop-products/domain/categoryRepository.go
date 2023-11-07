package domain

import "context"

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, id int) (Category, error)
	AddCategory(ctx context.Context, category Category) error
	UpdateCategory(ctx context.Context, category Category) error
	DeleteCategory(ctx context.Context, id int) error
}
