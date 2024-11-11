// POINTERS 

package main
import "fmt"

func main(){
	// var p *int32= new(int32)
	// var i int32
	// fmt.Printf("The value p points t is :%v", *p)
	// fmt.Printf("\nThe value if i is :%v",i)
	// changing value of memory location of a pointer
	// p = &i
	// *p = 1
	// fmt.Printf("\nThe value p points to is :%v",*p)
	// fmt.Printf("\nThe value if i is :%v",i)
	// var  k int32=6
	// i = k
	// fmt.Printf("\nThe value i points to is : %v",i)
	

	//WORKING WITH SLICES
	// var slice  = []int{1,2,4}
	// var sliceCopy = slice
	// sliceCopy[2]=4
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)


	// WORKING WITH FUNCTIONS 
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of thing1 is : %p", &thing1)
	var result [5]float64 =square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is :%v  ",thing1)


}

// Define the 'square' function that takes an array of 5 float64 elements as input
// and returns an array of 5 float64 elements

// Added pointers to save on memory and time 
func square(thing2 *[5]float64) [5]float64 {
    // Iterate over each element in the array
	fmt.Printf("\nThe memory location of the thing2 array is : %p", thing2)
    for i := range thing2 {
        // Square the value of each element
        thing2[i] = thing2[i] * thing2[i]
    }
    // Return the modified array
    return *thing2
}