package handler

import (
	"errors"
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/encoding"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/request"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/utils"
	commonErr "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/errors"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/models"
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

		if shortURL.ShortURL != "" {

			_, err = urlRepo.GetOriginalURL(shortURL.ShortURL)
			if err != nil {
				if errors.Is(err, commonErr.ErrInternalServerError) {
					HandleInternalServerError(w, r, err, logger, "unable to get original url")
					return
				} else {
					response := map[string]string{
						"short_url": shortURL.ShortURL,
						"message":   "short url already exists",
					}
					encoding.EncodeJson(w, r, http.StatusConflict, response)
					return
				}
			}
		} else {
			shortURL.ShortURL = utils.GenerateShortURL()
		}

		modelURL, err := urlRepo.CreateShortURL(&models.ShortURL{
			ShortURL:    shortURL.ShortURL,
			OriginalURL: shortURL.OriginalURL,
		})
		if err != nil {
			HandleInternalServerError(w, r, err, logger, "unable to create short url")
			return
		}

		response := request.ShortUrlResp{
			ID:          modelURL.ID,
			ShortURL:    modelURL.ShortURL,
			OriginalURL: modelURL.OriginalURL,
			CreatedAt:   modelURL.CreatedAt,
			UpdatedAt:   modelURL.UpdatedAt,
		}
		err = encoding.EncodeJson(w, r, http.StatusCreated, response)
		if err != nil {
			HandleInternalServerError(w, r, err, logger, "unable to create short url")
			return
		}

	}
}
