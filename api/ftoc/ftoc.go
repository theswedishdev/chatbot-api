package ftoc

import (
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	temperature := r.URL.Query().Get("temperature")
	
	if (temperature == "") {
		fmt.Fprintf(w, "Missing user input")
		return
	}
	
	temperatureInFahrenheit, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid input")
		return
	}
	
	temperatureInCelsius := (temperatureInFahrenheit * 1.8) + 32
	
	fmt.Fprintf(w, "%.2f°C", temperatureInCelsius)
}
