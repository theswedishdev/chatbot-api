package ctof

import (
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	temperature := r.URL.Query().Get("temperature")
	
	if (temperature == "") {
		return
	}
	
	temperatureInCelsius, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	
	temperatureInFahrenheit := (temperatureInCelsius * 1.8) + 32
	
	fmt.Fprintf(w, "%02f", temperatureInFahrenheit)
}
