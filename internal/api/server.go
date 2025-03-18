package main

import (
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/middleware"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/routes"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func NewServer(cfg *config.Config, logger *jsonlog.Logger, urlRepo repository.UrlShortener) http.Handler {
	mux := http.NewServeMux()

	routes.AddRoutes(mux, cfg, urlRepo)

	var handler http.Handler = mux
	handler = middleware.LoggingMiddleware(logger, handler)

	return handler
}
