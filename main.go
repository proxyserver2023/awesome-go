package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	// input
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -2, -3, -4, -5, -6, -7, -8, -9}

	// init channel
	c := make(chan int)

	// asking goroutines for using channel to send the total
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// receiving the data from channel
	x, y := <-c, <-c

	// output
	fmt.Println(x, y, x+y)
}
