package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string `json:"full_name"`
	Weapon string
	Level  int
}

func main() {
	sForm := `{"full_name": "suraj", "Weapon":"ninja kunai", "Level": 1}`
	fmt.Println(sForm)
	var suraj Ninja
	err := json.Unmarshal([]byte(sForm), &suraj)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		return
	}
	fmt.Println(suraj)
	sto, err := json.Marshal(suraj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", sto)
	fmt.Printf("%s\n", sto)

}
