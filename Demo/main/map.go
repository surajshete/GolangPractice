package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil)
	m = make(map[string]string)
	fmt.Println(m == nil)
	m["suraj"] = "ninja"
	m["suraj"] += " programmer"
	fmt.Println(m)
	for name, title := range m {
		fmt.Println(name + " " + title)
	}
	title,ok :=m["suraj1"]
	if ok{
		fmt.Println(title)
	}else{
		fmt.Println("Not found")
	}
}
