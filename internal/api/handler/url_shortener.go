package handler

import (
	"errors"
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/encoding"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/request"
	commonErr "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/errors"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func CreateShortLink(logger *jsonlog.Logger, urlRepo repository.UrlShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		shortURL, problems, err := encoding.Validated[*request.ShortURL](r)

		if len(problems) > 0 {
			err = encoding.EncodeJson(w, r, http.StatusBadRequest, problems)
			if err != nil {
				HandleInternalServerError(w, r, err, logger, "failed to encode validation errors")
			}
			return
		}
		if err != nil {
			HandleInternalServerError(w, r, err, logger, "failed to encode validation errors")
			return
		}

		if shortURL.ShortURL != "" {
			orginalURL, err := urlRepo.GetOriginalURL(shortURL.ShortURL)
			if err != nil {
				if !errors.Is(err, commonErr.ErrShortURLNotFound) {
					HandleInternalServerError(w, r, err, logger, "failder to get the original url from database")
				}
			}
			if orginalURL != "" {

			}
		}

	}

}
