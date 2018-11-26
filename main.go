package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"

	fmt.Println("map: => ", m)

	v1 := m["key1"]
	v2 := m["key2"]
	fmt.Println("len:", len(m), " v1: ", v1, " v2: ", v2)

	delete(m, "key2")
	fmt.Println("map: => ", m)

	_, ok := m["k2"]
	fmt.Println("m[k2] exists ? ", ok)

	n := map[string]int{
		"foo":    1,
		"bar":    2,
		"foobar": 3,
	}
	fmt.Println(n)
}
