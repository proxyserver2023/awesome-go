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

	// Pass by value
	// v := reflect.ValueOf(x)
	// v.SetFloat(7.1)

	// Pass by reference
	p := reflect.ValueOf(&x)
	vv := p.Elem()
	vv.SetFloat(7.1)
}
