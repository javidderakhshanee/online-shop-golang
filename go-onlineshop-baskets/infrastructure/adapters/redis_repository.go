package adapters

import (
	"context"
	"encoding/json"
	"onlineshopbasket/config"
	domain "onlineshopbasket/domain"

	"github.com/go-redis/redis/v8"
)

var configuration config.Configuration

type RedisRepository struct {
	db            *redis.Client
	ctx           context.Context
	configuration config.Configuration
}

func NewRedisRepository() *RedisRepository {
	configuration = config.NewConfiguration()

	return &RedisRepository{
		db:            getRedisClient(configuration.Redis),
		ctx:           context.Background(),
		configuration: configuration,
	}
}

func getRedisClient(redisConfig config.RedisConfigurations) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisConfig.ConnectionString,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
}
func getKey(id string) string {
	return configuration.Redis.InstanceName + "_" + id
}
func (repo *RedisRepository) GetBasket(ctx context.Context, id string) (domain.BasketHeader, error) {
	var basket domain.BasketHeader

	s, err := repo.db.Get(repo.ctx, getKey(id)).Result()

	json.Unmarshal([]byte(s), &basket)

	return basket, err
}

func (repo *RedisRepository) UpdateBasket(ctx context.Context, id string, b domain.BasketHeader) error {
	err := repo.db.Set(repo.ctx, getKey(id), b, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (repo *RedisRepository) DeleteBasket(ctx context.Context, id string) error {
	err := repo.db.Del(repo.ctx, getKey(id)).Err()
	if err != nil {
		return err
	}

	return nil
}
