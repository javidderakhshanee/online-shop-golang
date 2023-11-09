package app

import (
	"context"

	categorycommand "onlineshopproduct/app/commands/category"
	categoryquery "onlineshopproduct/app/queries/category"

	productcommand "onlineshopproduct/app/commands/product"
	productquery "onlineshopproduct/app/queries/product"

	"onlineshopproduct/infrastructure/adapters"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(ctx context.Context) Application {
	repoCategory := adapters.NewCategoryRepository(ctx)
	repoProduct := adapters.NewProductRepository(ctx)

	return Application{
		Commands: Commands{
			AddCategory:    categorycommand.NewAddCategoryHandler(repoCategory),
			UpdateCategory: categorycommand.NewUpdateCategoryHandler(repoCategory),
			DeleteCategory: categorycommand.NewDeleteCategoryHandler(repoCategory),

			AddProduct:    productcommand.NewAddProductHandler(repoProduct),
			UpdateProduct: productcommand.NewUpdateProductHandler(repoProduct),
			DeleteProduct: productcommand.NewDeleteProductHandler(repoProduct),
		},
		Queries: Queries{
			GetCategories: categoryquery.NewGetCategoriesHandler(repoCategory),
			GetCategory:   categoryquery.NewGetCategoryHandler(repoCategory),
			GetProducts:   productquery.NewGetProductsHandler(repoProduct),
			GetProduct:    productquery.NewGetProductHandler(repoProduct),
		},
	}
}

type Queries struct {
	GetCategory   categoryquery.GetCategoryHandler
	GetCategories categoryquery.GetCategoriesHandler

	GetProduct  productquery.GetProductHandler
	GetProducts productquery.GetProductsHandler
}

type Commands struct {
	AddCategory    categorycommand.AddCategoryHandler
	UpdateCategory categorycommand.UpdateCategoryHandler
	DeleteCategory categorycommand.DeleteCategoryHandler

	AddProduct    productcommand.AddProductHandler
	UpdateProduct productcommand.UpdateProductHandler
	DeleteProduct productcommand.DeleteProductHandler
}
