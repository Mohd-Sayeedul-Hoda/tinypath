package redis

import (
	"context"
	"errors"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/cache"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	commonErr "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/errors"
	"github.com/redis/go-redis/v9"
)

type cacheRepo struct {
	client *redis.Client
}

func NewCacheRepo(cfg *config.Config) (cache.CacheRepo, error) {
	opts, err := redis.ParseURL(cfg.Cache.DSN)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)

	return &cacheRepo{
		client: client,
	}, nil
}

func (c *cacheRepo) Get(key string) (string, error) {
	value, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", commonErr.ErrShortURLNotFound
		}
		return "", commonErr.NewCustomInternalErr(err)

	}
	return value, nil
}

func (c *cacheRepo) Set(key, value string) error {
	err := c.client.Set(context.Background(), key, value, time.Minute*30).Err()
	if err != nil {
		return commonErr.NewCustomInternalErr(err)
	}
	return nil
}

func (c *cacheRepo) Delete(key string) error {
	err := c.client.Del(context.Background(), key).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return commonErr.ErrShortURLNotFound
		}
		return commonErr.NewCustomInternalErr(err)
	}
	return nil
}
