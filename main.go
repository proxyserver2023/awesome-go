package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string, ii time.Duration) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		time.Sleep(ii * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	go say("world", 1000)
	say("Hello", 1000)
}
