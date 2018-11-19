package gopl

import (
	"bufio"
	"fmt"
	"os"
)

func DuplicateText() {
	// new scanner

	input := bufio.NewScanner(os.Stdin)
	m := make(map[string]int)
	for i := 0; input.Scan(); i++ {
		// fmt.Println(input.Text())
		m[input.Text()]++
		if i > 10 {
			break
		}
	}

	for i, mId := range m {
		if mId > 1 {
			fmt.Println(i, mId)
		}
	}

}
