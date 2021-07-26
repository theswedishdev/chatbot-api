package commands

import (
	"fmt"
	"strings"
)

func FToK(fahrenheit float64) string {
	result := fmt.Sprintf("%.1f°K", ConvertFahrenheitToKelvin(fahrenheit))
	return strings.Replace(result, ".0", "", -1)
}
