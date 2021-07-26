package commands

import (
	"fmt"
	"strings"
)

func KToF(kelvin float64) string {
	result := fmt.Sprintf("%.1f°F", ConvertFahrenheitToCelsius(kelvin))
	return strings.Replace(result, ".0", "", -1)
}
