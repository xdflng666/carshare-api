package getCarLocations

import (
	"carshare-api/internal/models"
	resp "carshare-api/lib/api/response"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=CarLocationsGetter
type CarLocationsGetter interface {
	GetCarLocations() ([]models.CarLocation, error)
}

func New(locationsGetter CarLocationsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getCarLocations.New"
		w.Header().Set("Access-Control-Allow-Origin", "*")

		res, err := locationsGetter.GetCarLocations()
		if err != nil {
			log.Printf("%s:%s", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Internal server error"))
		}

		render.JSON(w, r, res)
	}
}
