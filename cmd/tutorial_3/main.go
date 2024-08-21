package main 

import "fmt"

func main(){
	// declare a variable to store the string "Hello, World!"
	var printValue string ="Hello, World!"
	// call the function printMe with the variable printValue as an argument
	printMe(printValue)
}
// function to print "Hello, World!"
func printMe(printValue string){
	// print the value of the variable printValue
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int{
	// return the quotient of the numerator and denominator
	var result int = numerator / denominator
	return result

}