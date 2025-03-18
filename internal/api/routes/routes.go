package routes

import (
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/handler"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func AddRoutes(mux *http.ServeMux, cfg *config.Config, urlRepo repository.UrlShortener, loggger *jsonlog.Logger) {

	mux.Handle("/", handler.HandleRoot(loggger))

	mux.Handle("/v1/api/healthcheck", handler.HealthCheck(loggger))

}
