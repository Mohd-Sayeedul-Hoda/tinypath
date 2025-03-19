package routes

import (
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api/handler"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func AddRoutes(mux *http.ServeMux, cfg *config.Config, loggger *jsonlog.Logger, urlRepo repository.UrlShortener) {

	mux.HandleFunc("/", handler.HandleRoot(loggger))
	mux.HandleFunc("/api/v1/healthcheck", handler.HealthCheck(loggger))

	//tiny url paths
	mux.HandleFunc("/api/v1/short", handler.CreateShortLink(loggger, urlRepo))

}
