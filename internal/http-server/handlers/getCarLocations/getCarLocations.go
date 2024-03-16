package getCarLocations

import (
	"carshare-api/internal/models"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type CarLocationsGetter interface {
	GetCarLocations() ([]models.CarLocation, error)
}

func New(locationsGetter CarLocationsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		const op = "handlers.getCarLocations.New"

		res, err := locationsGetter.GetCarLocations()
		if err != nil {
			log.Fatal("Поешь говна")
		}

		render.JSON(w, r, res)

		//render.JSON(w, r, models.CarLocation{
		//	Name:      "Hi",
		//	UUID:      "1234",
		//	IsActive:  false,
		//	Lat:       1,
		//	Lon:       2,
		//	CreatedAt: time.Now(),
		//})
	}
}
