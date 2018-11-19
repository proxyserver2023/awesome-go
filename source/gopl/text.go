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

func DuplicateTextStdInAndFile() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			CheckErrPrint(err)
			countLines(f, counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
