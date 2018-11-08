package main

import (
	"fmt"
)

type One struct {
}

type Thinger interface {
	Hello()
	There()
}

func (o One) Hello() {
	fmt.Println("Hello, Playground")
}

func Run(t Thinger) {
	t.Hello()
}

func main() {
	t := One{}
	Run(t)
}
