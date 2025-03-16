package models

import "time"

type ShortURL struct {
	ID          string
	ShortURL    string
	OriginalURL string
	AccessCount int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Pagination struct {
	Limit  int
	OffSet int
}
