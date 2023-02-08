package main

import (
	"fmt"
	"os"
	"text/template"
)

type task struct {
	Name     string
	Language string
}

func main() {

	templatePath, err := os.Getwd()
	templatePath = templatePath + "/Template.txt"
	fmt.Println(templatePath)
	task := task{"suraj", "GO language"}

	t, err := template.New("Template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, task)
	if err != nil {
		fmt.Println(err)
	}
}
