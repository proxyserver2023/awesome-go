package main

import "fmt"

// app "github.com/alamin-mahamud/awesome-go/source/simpleapi"
// app "github.com/alamin-mahamud/awesome-go/source/gopl"
// app "github.com/alamin-mahamud/awesome-go/source/golang"

type T struct {
	id int
}

func main() {
	//app.Run()
	t1 := T{id: 1}
	t2 := T{id: 2}
	ts1 := []T{t1, t2}
	ts2 := []*T{}
	for _, t := range ts1 {
		ts2 = append(ts2, &t)
	}

	for _, t := range ts2 {
		fmt.Println((*t).id)
	}
}
