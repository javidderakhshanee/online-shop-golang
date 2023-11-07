package domain

import "context"

type ProductRepository interface {
	GetProducts(ctx context.Context, categoryId int) ([]Product, error)
	GetProduct(ctx context.Context, id int) (Product, error)
	AddProduct(ctx context.Context, product Product) error
	UpdateProduct(ctx context.Context, product Product) error
	DeleteProduct(ctx context.Context, id int) error
}
