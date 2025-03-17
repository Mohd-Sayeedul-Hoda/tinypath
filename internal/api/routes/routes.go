package routes

import (
	"net/http"

	"github.com/Mohd-Sayeedul-Hoda/"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
)

func addRoutes(mux http.ServeMux, cfg config) {

	mux.Handle("api/v1/shortlink")

}
