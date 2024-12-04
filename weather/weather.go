package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(geoData geo.GeoData, format int) string {
	urlStruct, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	urlStruct.RawQuery = params.Encode()

	resp, err := http.Get(urlStruct.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(body)
}
