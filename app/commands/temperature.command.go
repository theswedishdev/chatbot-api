package commands

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Convert a temperature from Fahrenheit to Celsius
func ConvertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Convert a temperature from Kelvin to Celsius
func ConvertKelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Convert a temperature from Celsius to Fahrenheit
func ConvertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// Convert a temperature from Kelvin to Fahrenheit
func ConvertKelvinToFahrenheit(kelvin float64) float64 {
	return ConvertCelsiusToFahrenheit(ConvertKelvinToCelsius(kelvin))
}

// Convert a temperature from Celsius to Kelvin
func ConvertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Convert a temperature from Fahrenheit to Kelvin
func ConvertFahrenheitToKelvin(fahrenheit float64) float64 {
	return ConvertCelsiusToKelvin(ConvertFahrenheitToCelsius(fahrenheit))
}

func GetTemperatureFloat(c *fiber.Ctx) (float64, error) {
	temperature := c.Params("temperature")
	temperature = strings.Replace(temperature, ",", ".", -1)

	temperatureFloat, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		return 0, fiber.NewError(fiber.ErrBadRequest.Code, "Invalid input")
	}

	return temperatureFloat, nil
}
