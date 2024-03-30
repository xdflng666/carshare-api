package postCarLocation

import (
	resp "carshare-api/lib/api/response"
	"errors"
	"github.com/go-chi/render"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	CarUUID string  `json:"car_uuid"`
}

type CarLocationPoster interface {
	PostCarLocation(lat, lon float64, carUUID string) error
}

func New(poster CarLocationPoster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.postCarLocation.New"

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом.
			// Обработаем её отдельно
			log.Println("request body is empty")
			render.JSON(w, r, resp.Error("empty request"))
			return
		}

		if err != nil {
			log.Println("failed to decode request body", err)
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Println(op, "request body decoded", req)

		err = poster.PostCarLocation(req.Lat, req.Lon, req.CarUUID)
		if err != nil {
			log.Println(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Internal server error"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
