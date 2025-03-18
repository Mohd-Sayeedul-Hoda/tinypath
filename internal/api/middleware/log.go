package middleware

import (
	"net/http"
	"time"

	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
)

func LoggingMiddleware(logger *jsonlog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		properties := map[string]string{
			"method":     r.Method,
			"uri":        r.RequestURI,
			"remoteAddr": r.RemoteAddr,
			"duration":   duration.String(),
		}
		logger.PrintInfo("HTTP Request", properties)
	})
}
