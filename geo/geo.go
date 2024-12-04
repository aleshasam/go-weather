package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetCurrentLocation(city string) (*GeoData, error) {
	if city != "" {
		if !checkCity(city) {
			panic(city)
		}
		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("Status code is not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}

	return &geo, nil
}

func checkCity(city string) bool {
	postData, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var checkCityData CityPopulationResponse
	err = json.Unmarshal(body, &checkCityData)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if checkCityData.Error == true {
		fmt.Println("Такого города не существует")
		return false
	}

	fmt.Println("Город найден")
	return true
}
