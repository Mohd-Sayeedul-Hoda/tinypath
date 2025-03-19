package request

import (
	"context"
	"net/url"
	"strings"
)

type ShortURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func (s *ShortURL) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	s.OriginalURL = strings.TrimSpace(s.ShortURL)
	if s.OriginalURL == "" {
		problems["original_url"] = "original url cannot be empty"
	} else {
		_, err := url.ParseRequestURI(s.OriginalURL)
		if err != nil {
			problems["original_url"] = "original url should be valid url"
		}
	}

	s.ShortURL = strings.TrimSpace(s.ShortURL)

	return problems
}
