package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age:    " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)

	list[0] = 1
	list[1] = "2 as a string"
	list[2] = Person{
		name: "alamin",
		age:  24,
	}

	for _, elem := range list {
		if element, ok := elem.(int); ok {
			fmt.Printf("%T=>%v\n", element, element)
		} else if element, ok := elem.(string); ok {
			fmt.Printf("%T=>%v\n", element, element)
		} else if element, ok := elem.(Person); ok {
			fmt.Printf("%T=>%v\n", element, element)
		}
	}

}
