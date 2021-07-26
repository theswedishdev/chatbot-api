package commands

import (
	"fmt"
	"strings"
)

func CToF(celsius float64) string {
	result := fmt.Sprintf("%.1f°F", ConvertCelsiusToFahrenheit(celsius))
	return strings.Replace(result, ".0", "", -1)
}
