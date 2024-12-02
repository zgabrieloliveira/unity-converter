package main

import "errors"

func MetersToFeet(meters float64) (float64, error) {

	if meters < 0 {
		return 0, errors.New("Invalid input, meters must be positive")
	}

	return meters * 3.28084, nil
}

func FeetToMeters(feet float64) (float64, error) {

	if feet < 0 {
		return 0, errors.New("Invalid input, feet must be positive")
	}

	return feet / 3.28084, nil
}
