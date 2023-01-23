package main

import "fmt"

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)
	arr1 := [...]int{1: 5, 4: 2}
	fmt.Println(arr1)
}
