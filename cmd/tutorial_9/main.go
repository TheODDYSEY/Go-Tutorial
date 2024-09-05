// CHANNELS
// They hold data
// thread safe (data races avoid)
// listen for data

package main

import "fmt"

func main(){
	// create the channel 
	var c = make(chan int)

	// start the go routine
	go process(c)
	
	for i:= range c{
		fmt.Println(i)
	}
}

func process(c chan int){
	defer close(c)
	for i:=0 ; i<5 ; i++{
		c<-i
	} 
	
}

