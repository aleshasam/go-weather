package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrWrongFormat = errors.New("WRONGFORMAT")

func Get(geoData geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}

	urlStruct, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	urlStruct.RawQuery = params.Encode()

	resp, err := http.Get(urlStruct.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(body), nil
}
