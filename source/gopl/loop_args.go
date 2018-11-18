package gopl

import (
	"fmt"
	"os"
	"strings"
)

func LoopArgs() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func LoopArgsWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func LoopArgsWithZero() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
