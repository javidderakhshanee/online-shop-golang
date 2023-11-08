package app

import (
	"context"

	categorycommand "onlineshopproduct/app/commands/category"
	categoryquery "onlineshopproduct/app/queries/category"
	"onlineshopproduct/infrastructure/adapters"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(ctx context.Context) Application {
	repoCategory := adapters.NewCategoryRepository(ctx)

	return Application{
		Commands: Commands{
			AddCategory: categorycommand.NewAddCategoryHandler(repoCategory),
		},
		Queries: Queries{
			GetCategories: categoryquery.NewGetCategoriesHandler(repoCategory),
			GetCategory:   categoryquery.NewGetCategoryHandler(repoCategory),
		},
	}
}

type Queries struct {
	//GetProduct  queries.GetBasketHandler
	//GetProducts queries.GetBasketHandler

	GetCategory   categoryquery.GetCategoryHandler
	GetCategories categoryquery.GetCategoriesHandler
}

type Commands struct {
	AddCategory    categorycommand.AddCategoryHandler
	UpdateCategory categorycommand.UpdateCategoryHandler
	DeleteCategory categorycommand.DeleteCategoryHandler

	//AddProduct    commands.UpdateBasketHandler
	//UpdateProduct commands.UpdateBasketHandler
	//DeleteProduct commands.DeleteBasketHandler
}
