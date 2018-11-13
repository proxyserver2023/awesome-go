package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	t := reflect.TypeOf(x)
	v := reflect.TypeOf(x)

	fmt.Println("type: ", t)
	fmt.Println("value: ", v)
	fmt.Println("kind: ", v.Kind() == reflect.Float64)
}
