package main

import "fmt"

func main() {
	array_fixed_size := [...]int{1, 2, 3}
	fmt.Println(array_fixed_size)
	 //array_fixed_size=[...]int{1,2}//------------ it will
	// when we declare array with fixed size we cannot reassign same array with diffrent size in that case we use slice
	sliceA := []int{1, 3, 4}
	fmt.Println(len(sliceA))
	sliceA = []int{1, 4, 4, 4, 4}
	fmt.Println(sliceA) 
}
