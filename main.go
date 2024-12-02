package main

import (
	"fmt"
	"time"
)

func main() {
	buildMenu()
}

func buildMenu() {
	var option int
	var value float64

	for option != 7 {

		fmt.Println("------------------------------------------------")
		fmt.Println("\t\tUNIT CONVERTER")
		fmt.Println("------------------------------------------------")

		fmt.Println("Please select an option:\n")
		fmt.Println("1. Meters to Feet")
		fmt.Println("2. Feet to Meters")
		fmt.Println("3. Kilograms to Pounds")
		fmt.Println("4. Pounds to Kilograms")
		fmt.Println("5. Celsius to Fahrenheit")
		fmt.Println("6. Fahrenheit to Celsius")
		fmt.Println("7 Exit\n")

		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Please enter a distance in meters:")
			fmt.Scanln(&value)
			getResult(MetersToFeet, value, "Meters", "Feet")

		case 2:
			fmt.Println("Please enter a distance in feet:")
			fmt.Scanln(&value)
			getResult(FeetToMeters, value, "Feet", "Meters")

		case 3:
			fmt.Println("Please enter a weight in kilograms:")
			fmt.Scanln(&value)
			getResult(KilogramsToPounds, value, "Kilograms", "Pounds")

		case 4:
			fmt.Println("Please enter a weight in pounds:")
			fmt.Scanln(&value)
			getResult(PoundsToKilograms, value, "Pounds", "Kilograms")

		case 5:
			fmt.Println("Please enter a temperature in Celsius:")
			fmt.Scanln(&value)
			result := CelsiusToFahrenheit(value)
			fmt.Println("Celsius" + ": " + fmt.Sprintf("%f", value))
			fmt.Println("Fahrenheit" + ": " + fmt.Sprintf("%f", result))

		case 6:
			fmt.Println("Please enter a temperature in Fahrenheit:")
			fmt.Scanln(&value)
			result := FahrenheitToCelsius(value)
			fmt.Println("Fahrenheit" + ": " + fmt.Sprintf("%f", value))
			fmt.Println("Celsius" + ": " + fmt.Sprintf("%f", result))

		case 7:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option")
		}

		time.Sleep(2 * time.Second)
	}

}

func getResult(funcCalled func(float64) (float64, error), value float64, inputUnit string, outputUnit string) {
	result, err := funcCalled(value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inputUnit + ": " + fmt.Sprintf("%f", value))
		fmt.Println(outputUnit + ": " + fmt.Sprintf("%f", result))
	}
}
