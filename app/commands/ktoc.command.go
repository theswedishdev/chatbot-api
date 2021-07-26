package commands

import (
	"fmt"
	"strings"
)

func KToC(kelvin float64) string {
	result := fmt.Sprintf("%.1f°C", ConvertKelvinToCelsius(kelvin))
	return strings.Replace(result, ".0", "", -1)
}
