package main

import "errors"

func KilogramsToPounds(kilograms float64) (float64, error) {

	if kilograms < 0 {
		return 0, errors.New("Invalid input, kilograms must be positive")
	}

	return kilograms * 2.20462, nil
}

func PoundsToKilograms(pounds float64) (float64, error) {

	if pounds < 0 {
		return 0, errors.New("Invalid input, pounds must be positive")
	}

	return pounds / 2.20462, nil
}
