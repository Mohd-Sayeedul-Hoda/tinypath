package request

import (
	"context"
	"net/url"
	"strings"
	"time"
)

type ShortURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type UpdateShortURLRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortUrlResp struct {
	ID          string    `json:"id,omitempty"`
	ShortURL    string    `json:"short_url,omitempty"`
	OriginalURL string    `json:"original_url,omitempty"`
	AccessCount int       `json:"access_count,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *ShortURL) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	OriginalURLValidate(problems, s.OriginalURL)
	ShortURLValidate(problems, s.ShortURL)
	return problems
}

func (s *UpdateShortURLRequest) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	OriginalURLValidate(problems, s.OriginalURL)
	return problems
}

func ShortURLValidate(problems map[string]string, shortURL string) {

	parsedURL := strings.TrimSpace(shortURL)
	if parsedURL != "" {
		if len(parsedURL) > 8 {
			problems["short_url"] = "short url should be less or equal to 8"
		}
	}
}

func OriginalURLValidate(problems map[string]string, OriginalURL string) {

	OriginalURL = strings.TrimSpace(OriginalURL)
	if OriginalURL == "" {
		problems["original_url"] = "original url cannot be empty"
	} else {
		parsedURL, err := url.Parse(OriginalURL)
		if err != nil {
			problems["original_url"] = "original url should be valid url"
			return
		}
		if parsedURL.Scheme == "" {
			problems["original_url"] = "original url should include a schema (e.g., http:// or https://)"
		}

		if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
			problems["original_url"] = "original url scheme should be either http or https"
		}
	}
}
