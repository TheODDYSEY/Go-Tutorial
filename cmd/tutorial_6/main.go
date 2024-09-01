// STRUCTS AND INTERFACES IN GO
package main
import "fmt"

type gasEngine struct {
	// Define the fields of the struct
	mpg uint8
	gallons uint8
}
type electricEngine struct {
	// Define the fields of the struct
	kWh uint8
	mpkwh uint8
}



// method for calculating miles left in tank
func (e gasEngine) milesLeft() uint8 {
	// Define a method for the struct
	return e.mpg * e.gallons
}

// method for calculating miles left in battery
func (e electricEngine) milesLeft() uint8 {
	// Define a method for the struct
	return e.kWh * e.mpkwh
}


type engine interface{
	milesLeft() uint8
}


func canMakeIt (e engine , miles uint8){
	if miles <=e.milesLeft(){
		fmt.Printf("You can make it with %v miles left\n",e.milesLeft()-miles)
	}else{
		fmt.Println("You need to fuel up ")
	}
}

func main (){

	// var myEngine gasEngine = gasEngine{ 30 ,10}
	var myEngine electricEngine = electricEngine{25,15}
	canMakeIt(myEngine,50)

	// anonymous struct
	// var myEngine2 = struct {
	// 	mpg uint8
	// 	gallons uint8
	// }{mpg: 30, gallons: 10}




}