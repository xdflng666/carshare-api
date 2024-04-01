package postCarLocation

import (
	resp "carshare-api/lib/api/response"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Lat     float64 `json:"lat" validate:"required,number,gte=-90,lte=90"`
	Lon     float64 `json:"lon" validate:"required,number,gte=-180,lte=180"`
	CarUUID string  `json:"car_uuid" validate:"required,uuid"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=CarLocationPoster
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

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

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
