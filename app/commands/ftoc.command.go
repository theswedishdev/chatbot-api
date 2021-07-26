package commands

import (
	"fmt"
	"strings"
)

func FToC(fahrenheit float64) string {
	result := fmt.Sprintf("%.1f°C", ConvertFahrenheitToCelsius(fahrenheit))
	return strings.Replace(result, ".0", "", -1)
}
