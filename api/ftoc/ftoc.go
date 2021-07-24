package ftoc

import (
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
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
	
	temperatureInCelsius := (temperatureInFahrenheit - 32) * 5 / 9
	
	fmt.Fprintf(w, "%.2f°C", temperatureInCelsius)
}
