package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"weather-apps/utility"
)

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Agent")
	// ======================

	w.WriteHeader(http.StatusOK)

	humidity := utility.GenerateRandomNumber(0, 100)
	temperature := utility.GenerateRandomNumber(-10, 50)
	wind := utility.GenerateRandomNumber(0, 20)
	rain := utility.GenerateRandomNumber(0, 250)

	response := map[string]interface{}{
		"humidity":     humidity,
		"temperature":  temperature,
		"wind":         wind,
		"rain":         rain,
		"last_checked": time.Now(),
	}

	log.Printf("type=%v method=%v path=%v", "[INFO]", r.Method, r.URL.Path)

	json.NewEncoder(w).Encode(response)
}
