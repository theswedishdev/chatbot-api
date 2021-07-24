package ctof

import (
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	temperature := r.URL.Query().Get("temperature")

	if temperature == "" {
		fmt.Fprintf(w, "Missing user input")
		return
	}

	temperatureInCelsius, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid input")
		return
	}

	temperatureInFahrenheit := (temperatureInCelsius * 1.8) + 32

	fmt.Fprintf(w, "%.2f°F", temperatureInFahrenheit)
}
