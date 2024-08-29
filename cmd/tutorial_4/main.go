// Arrays are fixed size,same type ,indexable and contiguous  in memory

package main

import "fmt"

func main(){

	// init array method 1
	// var intArr [3]int32
	// intArr[1] =123

	// init array method 2
	// var intArr [3]int32 = [3]int32{1,2,3}

	// init array method 3
	// intArr := [3]int32{1,2,3}

	// init array method 4
	// intArr := [...]int32{1,2,3}
	// fmt.Println(intArr)


	// show array and element of array 
	// fmt.Println(intArr)
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// memory address of array
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// SLICES ARE JUST ARRAYS WITH ADDITIONAL FUNCTIONALITY 
	// SLICES ARE NOT FIXED SIZE
	// SLICES ARE REFERENCE TYPE
	// SLICES ARE INDEXABLE
	// SLICES ARE CONTIGUOUS IN MEMORY

	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4,5,6}
	// print capacity 
	fmt.Printf("The length is %v with capacity %v\n",len(intSlice),cap(intSlice))
	fmt.Println(intSlice)

	// append element to slice
	intSlice = append(intSlice,7)
	// print capacity
	fmt.Printf("The length is %v with capacity %v\n",len(intSlice),cap(intSlice))
	fmt.Println(intSlice)

	// append using spread operator
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice,intSlice2...)
	fmt.Println(intSlice)
	// print capacity
	fmt.Printf("The length is %v with capacity %v\n",len(intSlice),cap(intSlice))

	// slice using make function
	intSlice3 := make([]int32,3,100)
	fmt.Println(intSlice3)



	// MAPS
	// MAPS ARE COLLECTION OF KEY VALUE PAIRS
	// MAPS ARE DYNAMICALLY SIZED
	// MAPS ARE REFERENCE TYPE
	// MAPS ARE NOT INDEXABLE
	

} 