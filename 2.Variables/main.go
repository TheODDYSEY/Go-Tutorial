package main

// package fmt is used to print to the console
import "fmt"

// package unicode/utf8 is used to work with strings
import "unicode/utf8"

func main ()  {
	// int stores numbers without decimal points
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	// float stores numbers with decimal points
	var floatNum float32 = 32767.99999
	fmt.Println(floatNum)

	// cannot perform operations on different types unless they are converted to the same type
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 * float32(intNum32)
	fmt.Println(result)

	// int division returns the quotient
	var intNum1 int = 10
	var intNum2 int = 3
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// string stores text
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	// len returns the length of a string (bytes)
	fmt.Println(len("test"))

	// utf8.RuneCountInString returns the number of runes in a string
	fmt.Println(utf8.RuneCountInString("test"))

	// rune stores a single unicode character
	var myRune rune = 'a'
	fmt.Println(myRune)

	// bool stores true or false
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// nil is the zero value for pointers, interfaces, maps, slices, channels, and function types
	var intNum3 rune 
	fmt.Println(intNum3)

	// another way to declare a variable
	myVar := "text"
	fmt.Println(myVar)

	// multiple variable declaration
	var1, var2, var3 :=  1, 2, 3
	fmt.Println(var1, var2, var3)

	// constants cannot be changed
	const myConst string = "const value"
	fmt.Println(myConst)




}