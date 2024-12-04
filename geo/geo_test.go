package geo_test

import (
	"demo/weather/geo"
	"errors"
	"testing"
)

func TestGetCurrentLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	got, err := geo.GetCurrentLocation(city)

	if err != nil {
		t.Error("Ошибка получения города", err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

func TestGetCurrentLocationNoCity(t *testing.T) {
	city := "NotLondon"
	_, err := geo.GetCurrentLocation(city)
	if !errors.Is(err, geo.ErrNoCity) {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
