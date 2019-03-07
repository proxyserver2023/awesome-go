package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
)

type A struct {
	name string
}

type B struct {
	age int
	A
}

func (b B) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.age, validation.Required),
		validation.Field(&b.A.name, validation.Required),
	)
}

func main() {
	b := B{
		// A:   A{name: "Hello"},
		age: 10,
	}
	err := b.Validate()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", b)
}
