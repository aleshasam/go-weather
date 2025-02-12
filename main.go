package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()
	geoData, err := geo.GetCurrentLocation(*city)
	if err != nil {
		panic(err.Error())
	}
	weatherData, err := weather.Get(*geoData, *format)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(weatherData)
}
