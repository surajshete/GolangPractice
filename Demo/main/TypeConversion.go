package main

import "fmt"
func main()  {
	type fahrenheit int
	type celsius int
	var c celsius=0  
	var f fahrenheit=32
	fmt.Print(c,f)
	var n int
	fmt.Println("enter a num")
	fmt.Scanf("%d",&n)
	c=celsius((n-32)*5/9);
	fmt.Print(c)

}