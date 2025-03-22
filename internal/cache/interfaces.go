package cache

type CacheRepo interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}
