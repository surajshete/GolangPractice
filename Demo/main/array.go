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
	str1 := [...]string{"s1", "s2", "s3"}
	fmt.Println(str1)
	for _, v := range str1 {
		fmt.Println(v)
	}

}
