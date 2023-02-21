package main

import (
	"fmt"
	"os"
	"text/template"
)

type user struct {
	Name     string
	Language string
}

func main() {

	templatePath, err := os.Getwd()
	templatePath = templatePath + "/Template.txt"
	fmt.Println(templatePath)
	user1 := user{"suraj", "GO language"}

	t, err := template.New("Template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, user1)
	if err != nil {
		fmt.Println(err)
	}
	user2 := user{"sam", "GO language"}
	err = t.Execute(os.Stdout, user2)
	if err != nil {
		fmt.Println(err)
	}
}
