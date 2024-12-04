package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"errors"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3

	result, err := weather.Get(geoData, format)

	if err != nil {
		t.Error("Ошибка получения погоды", err.Error())
	}

	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "Zero", format: 0},
	{name: "Minus format", format: -10},
}

func TestGetWrongFormat(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}

			_, err := weather.Get(geoData, testCase.format)

			if !errors.Is(err, weather.ErrWrongFormat) {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrWrongFormat, err)
			}
		})
	}
}
