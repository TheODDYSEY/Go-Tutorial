// GO ROUTINES 
// Running concurrency 


package main 

import  (
	"fmt"
	"math/rand"
	"time"
)


// Simulating a database 
var dbData = []string{"id1","id2","id3","id4","id5"}

// iterate/ call items in database and will take 5 seconds each 
func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		dbCall(i)
	} 
	fmt.Printf("\nTotal execution time :%v",time.Since(t0))

}

func dbCall(i int){
	// simulate DB call delay
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is :" ,dbData[i])

}