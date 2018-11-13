package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "Name:" + h.name + ", Age:" + strconv.Itoa(h.age) + " years, Contact:" + h.phone
}

func main() {
	bob := Human{
		name:  "Alamin",
		age:   24,
		phone: "+880 168 70 60 434",
	}

	fmt.Println("The Human is", bob)
}
