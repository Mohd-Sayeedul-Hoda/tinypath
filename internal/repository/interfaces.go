package repository

import "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/models"

type UrlShortener interface {
	CreateShortURL(shortURL *models.ShortURL) (*models.ShortURL, error)
	UpdateShortURL(shortURL string, originalURL string) error
	DeleteShortURL(shortURL string) error
	GetOriginalURL(shortURL string) (string, error)
	IncrementAccessCount(shortURL string) error
	GetShortURL(shortURL string) (*models.ShortURL, error)
	GetAllShortURL(pagination models.Pagination) ([]models.ShortURL, error)
}
