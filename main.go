package main

import (
	"fmt"
)

const (
	PRE_SIGNAL = iota
	POST_SIGNAL

	STATE_INIT
	STATE_RUNNING
	STATE_SHUTTING_DOWN
	STATE_TERMINATE
)

func main() {
	fmt.Println(PRE_SIGNAL)
	fmt.Println(POST_SIGNAL)
	fmt.Println(STATE_INIT)
	fmt.Println(STATE_RUNNING)
	fmt.Println(STATE_SHUTTING_DOWN)
	fmt.Println(STATE_TERMINATE)
}
