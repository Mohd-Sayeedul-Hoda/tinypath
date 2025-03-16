package models

import "time"

type ShortURL struct {
	ID          string
	ShortURL    string
	LongURL     string
	AccessCount int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
