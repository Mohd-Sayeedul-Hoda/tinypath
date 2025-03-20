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

		shortURL, problems, err := encoding.Validated[*request.ShortURL](r)

		if len(problems) > 0 {
			respondWithJSON(w, r, http.StatusBadRequest, problems, logger)
			return
		}

		if shortURL.ShortURL != "" {

			exists, err := urlRepo.GetOriginalURL(shortURL.ShortURL)

			if err != nil {
				if !errors.Is(err, commonErr.ErrShortURLNotFound) {
					HandleInternalServerError(w, r, err, logger, "unable to get original url")
					return
				}
			}

			if exists != "" {
				response := map[string]string{
					"message": commonErr.ErrShortURLAlreadyExists.Error(),
				}
				respondWithJSON(w, r, http.StatusConflict, response, logger)
				return
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

		respondWithJSON(w, r, http.StatusOK, response, logger)
	}
}

func GetShortLink(logger *jsonlog.Logger, urlRepo repository.UrlShortener) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.PathValue("short")
		if shortURL == "" {
			response := map[string]string{
				"message": "short url value should not be empty",
			}
			respondWithJSON(w, r, http.StatusBadRequest, response, logger)
			return
		}

		urlModel, err := urlRepo.GetShortURL(shortURL)
		if err != nil {
			if errors.Is(err, commonErr.ErrShortURLNotFound) {
				response := map[string]string{
					"message": err.Error(),
				}
				respondWithJSON(w, r, http.StatusNotFound, response, logger)
				return
			} else {
				HandleInternalServerError(w, r, err, logger, "")
				return
			}
		}

		response := request.ShortUrlResp{
			ID:          urlModel.ID,
			ShortURL:    urlModel.ShortURL,
			OriginalURL: urlModel.OriginalURL,
			AccessCount: urlModel.AccessCount,
			CreatedAt:   urlModel.CreatedAt,
			UpdatedAt:   urlModel.UpdatedAt,
		}

		respondWithJSON(w, r, http.StatusOK, response, logger)
	}
}
