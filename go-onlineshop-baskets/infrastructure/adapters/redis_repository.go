package adapters

import (
	"context"
	"encoding/json"
	"onlineshopbasket/config"
	domain "onlineshopbasket/domain"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	db            *redis.Client
	ctx           context.Context
	configuration config.Configuration
}

func NewRedisRepository() *RedisRepository {
	c := config.NewConfiguration()

	return &RedisRepository{
		db:            getRedisClient(c.Redis.ConnectionString),
		ctx:           context.Background(),
		configuration: c,
	}
}

func getRedisClient(connectionString string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (repo *RedisRepository) GetBasket(ctx context.Context, id string) (domain.BasketHeader, error) {
	var basket domain.BasketHeader

	s, err := repo.db.Get(repo.ctx, id).Result()
	json.Unmarshal([]byte(s), &basket)

	return basket, err
}

func (repo *RedisRepository) UpdateBasket(ctx context.Context, id string, b domain.BasketHeader) error {
	err := repo.db.Set(repo.ctx, id, b, 0)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (repo *RedisRepository) DeleteBasket(ctx context.Context, id string, b domain.BasketHeader) error {
	err := repo.db.Set(repo.ctx, id, b, 0)
	if err != nil {
		return err.Err()
	}

	return nil
}
