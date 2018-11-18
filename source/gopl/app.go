package gopl

import (
	"fmt"
	"time"
)

func timeIt(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
}

func Run() {
	// timeIt(HelloWorld)
	// timeIt(LoopArgs)
	// timeIt(LoopArgsWithJoin)
	// timeIt(LoopArgsWithZero)
	// timeIt(LoopArgsWithIndex)
	timeIt(DuplicateText)
}
