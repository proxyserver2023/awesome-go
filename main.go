package main

import (
	"fmt"
)

func main() {
	a := make(map[int]string)
	a[1] = "Hello"
	a[2] = "World"
	a[3] = "Nawaz"
	a[4] = "Varun"
	a[5] = "Huma"

	val, ok := a[6]
	fmt.Println(val, ok)
}
