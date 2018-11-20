package main

import "fmt"

type T struct {
	id int
}

func main() {
	t1 := T{id: 1}
	t2 := T{id: 2}
	ts1 := []T{t1, t2}
	//ts2 := []T{T{id: 3}, T{id: 4}}
	ts3 := []*T{}

	for _, t := range ts1 {
		ts3 = append(ts3, &t)
	}

	for _, t := range ts3 {
		fmt.Println((*t).id)
	}
}
