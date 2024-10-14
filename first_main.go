package main

import (
	"errors"
	"fmt"
)

func firstMain() {
	result, remainder, err := getTheRemainderAndDivision(90, 0) // provide numerator and denominator

	// If there's no error, print the result and remainder
	if err == nil {
		fmt.Printf("The Result is %v and the remainder is %v\n", result, remainder)
	} else {
		fmt.Println("Error:", err.Error()) // If there's an error, print the error message
	}
}

func getTheRemainderAndDivision(numerator int, denominator int) (int, int, error) {

	// Check if the denominator is 0
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	// Calculate result and remainder
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
