package app

import (
	"context"

	category "onlineshopproduct/app/commands/category"
	"onlineshopproduct/infrastructure/adapters"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(ctx context.Context) Application {
	repoCategory := adapters.NewCategoryRepository()

	return Application{
		Commands: Commands{
			AddCategory: category.NewAddCategoryHandler(repoCategory),
		},
		Queries: Queries{},
	}
}

type Queries struct {
	//GetProduct  queries.GetBasketHandler
	//GetProducts queries.GetBasketHandler

	//GetCategory   queries.GetBasketHandler
	//GetCategories queries.GetBasketHandler
}

type Commands struct {
	AddCategory category.AddCategoryHandler
	//UpdateCategory commands.UpdateBasketHandler
	//DeleteCategory commands.DeleteBasketHandler

	//AddProduct    commands.UpdateBasketHandler
	//UpdateProduct commands.UpdateBasketHandler
	//DeleteProduct commands.DeleteBasketHandler
}
