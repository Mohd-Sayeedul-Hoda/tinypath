package redis

import (
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/cache"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	commonErr "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/errors"
)

type noCacheRepo struct{}

func NewNoCacheRepo(cfg *config.Config) cache.CacheRepo {
	return &noCacheRepo{}
}

func (n noCacheRepo) Set(key, value string) error {
	return nil
}

func (n noCacheRepo) Get(key string) (string, error) {
	return "", commonErr.ErrCacheUnavailabe
}

func (n noCacheRepo) Delete(key string) error {
	return commonErr.ErrCacheUnavailabe
}
