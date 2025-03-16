package postgres

import (
	"context"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/models"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepo struct {
	pool *pgxpool.Pool
}

func NewURLShortenerRepo(pool *pgxpool.Pool) repository.UrlShortener {
	return &URLRepo{
		pool: pool,
	}
}

func (u *URLRepo) CreateShortURL(urlInfo *models.ShortURL) (*models.ShortURL, error) {

	query := `INSERT INTO urls (original_url, short_url, access_count, created_at) VALUE ($1, $2, $3, $4) RETURING id`

	err := u.pool.QueryRow(context.Background(), query,
		urlInfo.OriginalURL,
		urlInfo.ShortURL,
		urlInfo.AccessCount,
		urlInfo.CreatedAt,
	).Scan(&urlInfo.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *URLRepo) UpdateShortURL(shortURL string, originalURL string) error {

	query := `UPDATE urls SET original_url = $1, update_at = NOW() 
	WHERE  short_url = $2`

	_, err := u.pool.Exec(context.Background(), query, originalURL, shortURL)
	return err
}

func (u *URLRepo) DeleteShortURL(shortURL string) error {

	query := `DELETE from urls where short_url = $1`
	_, err := u.pool.Exec(context.Background(), query, shortURL)
	return err
}

func (u *URLRepo) IncrementAccessCount(shortURL string) error {
	query := `UPDATE urls SET access_count = access_count + 1,
	update_at = NOW() where short_url = $1`

	_, err := u.pool.Exec(context.Background(), query, shortURL)
	return err
}

func (u *URLRepo) GetShortURLStats(shortURL string) (*models.ShortURL, error) {

	var urlInfo *models.ShortURL
	query := `SELECT original_url, short_url, access_count, created_at, updated_at FROM urls WHERE short_url = $1`

	err := u.pool.QueryRow(context.Background(), query, shortURL).Scan(&urlInfo)
	if err != nil {
		return nil, err
	}

	return urlInfo, nil
}

func (u *URLRepo) GetAllShortURL(pagination models.Pagination) ([]models.ShortURL, error) {

	query := `SELECT 
	original_url, short_url, access_count, created_at, updated_at 
	FROM urls 
	ORDER BY created_at DESC
  LIMIT $1 OFFSET $2
`

	rows, err := u.pool.Query(context.Background(), query, pagination.Limit, pagination.OffSet)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var urlModels []models.ShortURL

	for rows.Next() {
		var urlModel models.ShortURL
		if err := rows.Scan(&urlModel.ShortURL, &urlModel.OriginalURL, &urlModel.AccessCount, &urlModel.CreatedAt, &urlModel.UpdatedAt); err != nil {
			return nil, err
		}
		urlModels = append(urlModels, urlModel)
	}

	return urlModels, nil
}

func (u *URLRepo) GetOriginalURL(shortURL string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_url = $1`

	var originalURL string
	err := u.pool.QueryRow(context.Background(), query, shortURL).Scan(&originalURL)
	return originalURL, err
}
