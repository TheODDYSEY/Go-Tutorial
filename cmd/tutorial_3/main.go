package main

import (
	"errors"
	"fmt"
)

func main() {
	// declare a variable to store the string "Hello, World!"
	var printValue string = "Hello, World!"
	// call the function printMe with the variable printValue as an argument
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	// call the function intDivision with the variables numerator and denominator as arguments
	var result, remainder ,err = intDivision(numerator, denominator)
	if err != nil {}
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)

}

// function to print "Hello, World!"
func printMe(printValue string) {
	// print the value of the variable printValue
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int,error) {

	// declare a variable to store the error
	var err error

	if denominator == 0 {
		// return an error if the denominator is 0
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	}
	// return the quotient of the numerator and denominator
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err

}
