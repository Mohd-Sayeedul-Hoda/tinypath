package repository

import "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/models"

type UrlShorener interface {
	CreateShortURL(shortURL *models.ShortURL) (*models.ShortURL, error)
	UpdateShortURL(shortURL string, originalURL string) error
	DeleteShortURL(shortURL string) error
	GetOriginalURL(shortURL string) (string, error)
	IncrementAccessCount(shortURL string) error
	GetShortURLStats(shortURL string) (*models.ShortURL, error)
	GetAllShortURL() ([]models.ShortURL, error)
}
