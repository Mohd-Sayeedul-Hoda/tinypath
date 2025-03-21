package routes

import (
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/handler"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func AddRoutes(mux *http.ServeMux, cfg *config.Config, loggger *jsonlog.Logger, urlRepo repository.UrlShortener) {

	mux.HandleFunc("GET /", handler.HandleRoot(loggger))
	mux.HandleFunc("GET /api/v1/healthcheck", handler.HealthCheck(loggger))

	// redirect of short url
	mux.HandleFunc("GET /{short}", handler.ShortURLRedirect(loggger, urlRepo))

	//tiny url paths
	mux.HandleFunc("POST /api/v1/short", handler.CreateShortLink(loggger, urlRepo))
	mux.HandleFunc("GET /api/v1/short/{short}", handler.GetShortLink(loggger, urlRepo))
	mux.HandleFunc("DELETE /api/v1/short/{short}", handler.DeleteShortLink(loggger, urlRepo))
	mux.HandleFunc("PATCH /api/v1/short/{short}", handler.UpdateShortLink(loggger, urlRepo))

}
