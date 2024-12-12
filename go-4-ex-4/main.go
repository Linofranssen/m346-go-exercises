package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
// TODO: implement the function convertFahrenheitToCelsius

type Fahrenheit float64
type Celsius float64

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func (f Fahrenheit) convertFahrenheitToCelsiusWithStruct() float64 {
	return (float64(f) - 32) * 5 / 9
}

func (c Celsius) convertCelsiusToFahrenheitWithStruct() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	c1, c2, c3 := -1.0, 32.5, 48.3
	fmt.Printf("%v°C => %v°F\n", c1, convertCelsiusToFahrenheit(c1))
	fmt.Printf("%v°C => %v°F\n", c2, convertCelsiusToFahrenheit(c2))
	fmt.Printf("%v°C => %v°F\n", c3, convertCelsiusToFahrenheit(c3))

	// TODO: call the function convertFahrenheitToCelsius
	f1, f2, f3 := 1.0, 48.3, 100.0
	fmt.Printf("%v°F => %v°C\n", f1, convertFahrenheitToCelsius(f1))
	fmt.Printf("%v°F => %v°C\n", f2, convertFahrenheitToCelsius(f2))
	fmt.Printf("%v°F => %v°C\n", f3, convertFahrenheitToCelsius(f3))

	var cold Fahrenheit = 15.3
	var cozy Celsius = 23.0

	fmt.Printf("Cozy: %v°C => %v°F\n", cozy, cozy.convertCelsiusToFahrenheitWithStruct())
	fmt.Printf("Cold: %v°F => %v°C\n", cold, cold.convertFahrenheitToCelsiusWithStruct())
}
