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
	for input.Scan() {
		m[input.Text()]++
	}

	for _, mId := range m {
		if mId > 1 {
			fmt.Println(mId)
		}
	}

}
