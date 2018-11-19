package main

import "fmt"

type X struct {
	name string
	next *X
}

func main() {
	//x := X{name: "foo", next: &X{name: "bar"}}
	x := X{name: "alamin", next: &X{name: "ashraf"}}
	fmt.Println(x.name)
	fmt.Println(x.next.name)
	fmt.Println(x.next.next)
}
