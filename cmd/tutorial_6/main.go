// STRUCTS AND INTERFACES IN GO
package main
import "fmt"

type gasEngine struct {
	// Define the fields of the struct
	mpg uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func main (){
	var myEngine gasEngine = gasEngine{25 ,10,owner{"Alex"}}
	fmt.Println(myEngine.mpg , myEngine.gallons , myEngine.name)

}