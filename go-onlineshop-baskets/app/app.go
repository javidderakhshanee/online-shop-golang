package app

import (
	"context"
	"onlineshopbasket/app/commands"
	"onlineshopbasket/app/queries"
	"onlineshopbasket/infrastructure/adapters"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(ctx context.Context) Application {
	repo := adapters.NewRedisRepository()

	return Application{
		Commands: Commands{
			UpdateBasket: commands.NewUpdateBasketHandler(repo),
			AddToBasket:  commands.NewAddToBasketHandler(repo),
			DeleteBasket: commands.NewDeleteBasketHandler(repo),
		},
		Queries: Queries{
			GetBasket: queries.NewGetBasketHandler(repo),
		},
	}
}

type Queries struct {
	GetBasket queries.GetBasketHandler
}

type Commands struct {
	UpdateBasket commands.UpdateBasketHandler
	AddToBasket  commands.AddToBasketHandler
	DeleteBasket commands.DeleteBasketHandler
}
