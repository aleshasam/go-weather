package main

import (
	"demo/weather/geo"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	//format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetCurrentLocation(*city)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(geoData)
}
