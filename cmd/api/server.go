package main

import (
	"net/http"

	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func NewServer(cfg config, logger *jsonlog.Logger, urlRepo repository.UrlShortener) http.Handler {
	mux := http.NewServeMux()
}

func (app *application) serve() error {
	//srv := &http.Server{}
	return nil

}
