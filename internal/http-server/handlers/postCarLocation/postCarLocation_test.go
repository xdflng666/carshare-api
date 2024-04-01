package postCarLocation_test

import (
	"bytes"
	"carshare-api/internal/http-server/handlers/postCarLocation"
	"carshare-api/internal/http-server/handlers/postCarLocation/mocks"
	"carshare-api/lib/api/response"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostCarLocation(t *testing.T) {
	cases := []struct {
		name      string
		lat       float64
		lon       float64
		UUID      string
		respError string
		mockError error
	}{
		{
			name: "Success",
			lat:  59.938088,
			lon:  30.313504,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
		},
		{
			name: "Empty UUID",
			lat: 1,
			lon: 1,
			respError: "field CarUUID is a required field",
		},
		{
			name: "Invalid UUID",
			lat: 1,
			lon: 1,
			UUID: "invalid uuid here...",
			respError: "field CarUUID is not valid",
		},
		{
			name: "Empty latitude",
			lon: 1,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
			respError: "field Lat is a required field",
		},
		{
			name: "Invalid latitude(too big)",
			lat: 300,
			lon: 1,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
			respError: "field Lat is not valid",
		},
		{
			name: "Invalid latitude(too small)",
			lat: -300,
			lon: 1,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
			respError: "field Lat is not valid",
		},
		{
			name: "Empty longitude",
			lat: 1,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
			respError: "field Lon is a required field",
		},
		{
			name: "Invalid longitude(too big)",
			lat: 1,
			lon: 666,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
			respError: "field Lon is not valid",
		},
		{
			name: "Invalid longitude(too small)",
			lat: 1,
			lon: -666,
			UUID: "550e8400-e29b-41d4-a716-446655440000",
			respError: "field Lon is not valid",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			carLocationPosterMock := mocks.NewCarLocationPoster(t)

			if tc.respError == "" || tc.mockError != nil {
				carLocationPosterMock.On("PostCarLocation", tc.lat, tc.lon, tc.UUID).Return(tc.mockError).Once()
			}

			handler := postCarLocation.New(carLocationPosterMock)

			input := fmt.Sprintf(`{"lat": %v, "lon": %v, "car_uuid": "%v"}`, tc.lat, tc.lon, tc.UUID)

			req, err := http.NewRequest(http.MethodPost, "/api/postLocation", bytes.NewReader([]byte(input)))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, rr.Code, http.StatusOK)

			body := rr.Body.String()

			// fmt.Println(body)

			var resp response.Response

			require.NoError(t, json.Unmarshal([]byte(body), &resp))

			require.Equal(t, tc.respError, resp.Error)

		})
	}

}
