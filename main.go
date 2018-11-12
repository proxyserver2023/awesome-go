package main

import "fmt"

func main() {

	// cmd.Execute()

	amy := struct {
		name string
		age  int
	}{"amy", 25}
	fmt.Println(amy.name)
}
