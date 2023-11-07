package domain

import "context"

type Repository interface {
	GetBasket(ctx context.Context, id string) (BasketHeader, error)
	UpdateBasket(ctx context.Context, id string, basket BasketHeader) error
	DeleteBasket(ctx context.Context, id string) error
}
