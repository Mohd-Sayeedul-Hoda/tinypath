package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type URLShortenerRepo struct {
	pool *pgxpool.Pool
}

func NewURLShortenerRepo(pool *pgxpool.Pool) *URLShortenerRepo {
	return &URLShortenerRepo{
		pool: pool,
	}
}
