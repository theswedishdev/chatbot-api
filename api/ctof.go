package ctof

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	temperature := r.URL.Query().Get("temperature")
	
	if (temperature == "") {
		fmt.Fprintf(w, "Missing user input")
		return
	}
	
	temperatureInCelsius, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid input")
		return
	}
	
	temperatureInFahrenheit := math.Round(temperatureInCelsius * 1.8) + 32
	
	fmt.Fprintf(w, "%02f°F", temperatureInFahrenheit)
}
