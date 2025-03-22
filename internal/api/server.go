package api

import (
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/middleware"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/routes"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/cache"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func NewServer(cfg *config.Config, logger *jsonlog.Logger, urlRepo repository.UrlShortener, cacheRepo cache.CacheRepo) http.Handler {
	mux := http.NewServeMux()

	routes.AddRoutes(mux, cfg, logger, urlRepo, cacheRepo)

	var handler http.Handler = mux
	handler = middleware.LoggingMiddleware(logger, handler)

	return handler
}
