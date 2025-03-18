package middleware

import (
	"net/http"
	"strconv"
	"time"

	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger *jsonlog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		startTime := time.Now()

		next.ServeHTTP(lrw, r)
		duration := time.Since(startTime)

		properties := map[string]string{
			"status_code": strconv.Itoa(lrw.statusCode),
			"status":      http.StatusText(lrw.statusCode),
			"method":      r.Method,
			"uri":         r.RequestURI,
			"remoteAddr":  r.RemoteAddr,
			"duration":    duration.String(),
		}
		logger.PrintInfo("HTTP Request", properties)
	})
}
