package main

import (
	"log"
	"net/http"
	"weather-apps/app"
)

func main() {
	port := ":9999"

	http.HandleFunc("/weather", app.GetCurrentWeather)
	log.Println("server running at port", port)

	http.ListenAndServe(port, nil)
}
