// STRING RUNES AND BYTES

package main

import ("fmt"
		"strings")

func main() {

	// METHOD 1
    // Declare a string variable and initialize it with the value "resume"
    // var myString = "resume"
    
    // Access the first character of the string (index 0) and store it in the variable 'indexed'
    // In Go, strings are indexed by bytes, so 'indexed' will be a byte (uint8)
    // var indexed = myString[0]
    
    // Print the value and type of the 'indexed' variable
    // %v is used to print the value, %T is used to print the type
    // fmt.Printf("%v ,%T\n", indexed, indexed)
    
    // Iterate over the string using a for loop with range
    // 'i' will be the index of the character, 'v' will be the Unicode code point (rune) of the character
    // fmt.Println("Characters in the string:")
    // for i, v := range myString {
        // Print the index and the Unicode code point of each character in the string
    //     fmt.Printf("Index: %d, Character: %c, Unicode: %U\n", i, v, v)
    // }

    // Convert the string to a byte slice
    // byteSlice := []byte(myString)
    
    // Print the byte slice
    // fmt.Println("String in bytes:")
    // for i, b := range byteSlice {
        // Print the index and the byte value of each character in the string
    //     fmt.Printf("Index: %d, Byte: %d\n", i, b)
    // }


	// METHOD 2
	// Declare a string variable and initialize it with the value "resume"
	// Convert the string to a slice of runes to handle Unicode characters properly
	var myString = []rune("resume")

	// Access the second character of the rune slice (index 1) and store it in the variable 'indexed'
	// In Go, runes are used to represent Unicode code points
	var indexed = myString[1]

	// Print the value and type of the 'indexed' variable
	// %v is used to print the value, %T is used to print the type
	fmt.Printf("%v ,%T\n", indexed, indexed)

	// Iterate over the rune slice using a for loop with range
	// 'i' will be the index of the character, 'v' will be the Unicode code point (rune) of the character
	for i, v := range myString {
		// Print the index and the Unicode code point of each character in the rune slice
		fmt.Println(i, v)
	}

	// Print the length of the rune slice
	// len(myString) returns the number of runes in the slice
	fmt.Printf("\nThe length of 'myString' is %d\n", len(myString))


	// Declaration of a rune type variable
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v",myRune)



	// METHOD 1
	// Concatenating strings using runes 
	// var strSlice = []string {"E","L","I","X","I","R"}
	// var catStr = ""
	// for i := range strSlice{
	// 	catStr += strSlice[i]
	// }
	// fmt.Printf("\n%s",catStr)
	// strIngs are immutable in go


	// METHOD 2 
	// Concatenating strings using STRING BUILDER
	var strSlice = []string {"E","L","I","X","I","R"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%s",catStr)




}