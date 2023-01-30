package main

import "fmt"

func main() {
	type Ninja struct {
		name   string
		weapon []string
		level  int
	}
	suraj := Ninja{name: "suraj", weapon: []string{"ninja star", "ninja kunai"}, level: 1}
	fmt.Print(suraj)
	vijay := Ninja{"vijay", []string{"star"}, 2}
vijay.level++;
	fmt.Println(vijay)
	type Dojo struct {
		ninja Ninja
		role  string
	}
	type DojoByReffrence struct{
		ninja *Ninja
		name string
	}
	Dojo_ninja := Dojo{suraj, "HOD"} //pass by value
	Dojo_ninja2 := DojoByReffrence{&suraj, "HOD"} //pass by value
	fmt.Println(Dojo_ninja)
	fmt.Println(Dojo_ninja2.ninja) //it will give values with &
	fmt.Println(*Dojo_ninja2.ninja)//it will give actual values
	fmt.Println(Dojo_ninja2.ninja.name)

	ninjaPointer :=new(Ninja)
	fmt.Println(ninjaPointer)
	fmt.Println(*ninjaPointer)
}
