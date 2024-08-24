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

	// if else statement to check if there is an error

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// }else{
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }

	// switch statement to check if there is an error
	// no need for writing the block keyword
	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)		
	}

	// switch statement to check the remainder
	switch remainder{
		case 0:
			fmt.Printf("The division was exact\n")	
		case 1,2:
			fmt.Printf("The division was close\n")
		default:
			fmt.Printf("The division was not close\n")		
	}


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
