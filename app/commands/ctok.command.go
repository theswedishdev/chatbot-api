package commands

import (
	"fmt"
	"strings"
)

func CToK(celsius float64) string {
	result := fmt.Sprintf("%.1f°K", ConvertCelsiusToKelvin(celsius))
	return strings.Replace(result, ".0", "", -1)
}
