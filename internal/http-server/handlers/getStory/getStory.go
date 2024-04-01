package getStory

import (
	"carshare-api/internal/models"
	resp "carshare-api/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=StoryGetter
type StoryGetter interface {
	GetStory(carUUID string) ([]models.Point, error)
}

func New(sg StoryGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getStory.New"

		carUUID := chi.URLParam(r, "car_uuid")

		res, err := sg.GetStory(carUUID)
		if err != nil {
			log.Printf("%s:%s", op, err)
			render.JSON(w, r, resp.Error("Internal server error"))
		}

		render.JSON(w, r, res)
	}
}
