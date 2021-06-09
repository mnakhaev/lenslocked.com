package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name         string
		City         string
		Age          int
		TaskToResult map[string]string
	}{
		"John Smith",
		"Moscow",
		21,
		map[string]string{
			"task1": "done",
			"task2": "not started",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
