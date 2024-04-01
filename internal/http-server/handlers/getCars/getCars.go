package getCars

import (
	"carshare-api/internal/models"
	resp "carshare-api/lib/api/response"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=CarsGetter
type CarsGetter interface {
	GetCars() ([]models.Car, error)
}

func New(cg CarsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getCars.New"
		w.Header().Set("Access-Control-Allow-Origin", "*")

		res, err := cg.GetCars()
		if err != nil {
			log.Printf("%s:%s", op, err)
			//w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Internal server error"))
		}

		render.JSON(w, r, res)
	}
}
