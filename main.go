package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) printLegs() int {
	fmt.Printf("%T => %v\n", c, c)
	return 4
}

type OctoCat struct {
	Cat
}

func (c Cat) legs() int {
	return 4
}

func (o OctoCat) legs() int {
	return 5
}

func main() {
	var o = OctoCat{Cat: Cat{name: "Hello"}}
	var c = Cat{name: "World"}

	// Octocat
	fmt.Println(o.printLegs())       // main.Cat -> {Hello} // 4
	fmt.Println(o.legs())            // 4
	fmt.Println("o.cat =>", o.Cat)   // {Hello}
	fmt.Println("o.name =>", o.name) // Hello

	// Cat
	fmt.Println(c.printLegs())       // main.Cat => {World} // 4
	fmt.Println(c.legs())            // 4
	fmt.Println("c.name =>", c.name) // World
}
