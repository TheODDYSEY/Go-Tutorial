// GO ROUTINES 
// Running concurrency 


package main 

import  (
	"fmt"
	"time"
	"sync"
)

// added wait groups 
var wg = sync.WaitGroup{}

// Simulating a database 
var dbData = []string{"id1","id2","id3","id4","id5"}

// slice to store all results from the database 
var results = []string{}

// iterate/ call items in database and will take 5 seconds each 
func main(){
	t0 := time.Now()
	for i:=0; i<1000; i++{
		// counts 
		wg.Add(1)
		go count()
	} 
	// wait for counter o go back to zero and then the code will execute 
	wg.Wait()
	fmt.Printf("\nTotal execution time :%v",time.Since(t0))
	fmt.Printf("\nThe results are %v",results)

}

// func dbCall(i int){
// 	// simulate DB call delay
// 	// removed randomness
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay)*time.Millisecond)
// 	wg.Done()

// }

func count(){
	var res int
	for i:=0; i<1000; i++{
		// counts 
		res+=1
	} 
	wg.Done()

}