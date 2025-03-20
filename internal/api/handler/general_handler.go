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

		err := encoding.EncodeJson(w, r, http.StatusOK, response)
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

		response := map[string]string{
			"status":  "ok",
			"message": "service is healthy",
		}

		err := encoding.EncodeJson(w, r, http.StatusOK, response)
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

func HandleInternalServerError(w http.ResponseWriter, r *http.Request, err error, logger *jsonlog.Logger, message string) {

	logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
		"error_message":  message,
	})

	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request, allowedMethods []string) {
// 	w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// }
