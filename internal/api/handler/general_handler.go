package handler

import (
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/encoding"
	"net/http"

	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
)

func HandleRoot(logger *jsonlog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		response := map[string]string{
			"message": "Welcome to Tiny URL!",
		}

		respondWithJSON(w, r, http.StatusOK, response, logger)
	}
}

func HealthCheck(logger *jsonlog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response := map[string]string{
			"status":  "ok",
			"message": "service is healthy",
		}

		respondWithJSON(w, r, http.StatusOK, response, logger)
	}

}

func HandleInternalServerError(w http.ResponseWriter, r *http.Request, err error, logger *jsonlog.Logger, message string) {

	logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
		"error_message":  message,
	})

	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func respondWithJSON[T any](w http.ResponseWriter, r *http.Request, status int, data T, logger *jsonlog.Logger) {

	err := encoding.EncodeJson(w, r, status, data)
	if err != nil {
		HandleInternalServerError(w, r, err, logger, "failed to encode JSON response")
	}
}
