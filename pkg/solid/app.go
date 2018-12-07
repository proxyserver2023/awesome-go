package solid

import "fmt"

// A is the base type in this package
type A struct {
	year int
}

// Greet greets us with a string
func (a A) Greet() { fmt.Println("Hello", a.year) }

// B is another type composed with A
type B struct {
	A
}

// Greet Overriding it
func (b B) Greet() { fmt.Println("Welcome", b.year) }

func Run() {
	var a A
	a.year = 2016

	var b B
	b.year = 2018

	a.Greet()
	b.Greet()
}
