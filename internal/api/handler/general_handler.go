package handler

import (
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/encoding"
	"net/http"

	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
)

func HandleRoot(logger *jsonlog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := map[string]string{
			"message": "Welcome to Tiny URL!",
		}

		err := encoding.WriteJson(w, http.StatusOK, r.Header, response)
		if err != nil {
			logger.PrintError(err, map[string]string{
				"request_method": r.Method,
				"request_url":    r.URL.String(),
			})
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func HealthCheck(logger *jsonlog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := map[string]string{
			"status":  "ok",
			"message": "service is healthy",
		}

		err := encoding.WriteJson(w, http.StatusOK, r.Header, response)
		if err != nil {
			logger.PrintError(err, map[string]string{
				"request_method": r.Method,
				"request_url":    r.URL.String(),
			})
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

}
