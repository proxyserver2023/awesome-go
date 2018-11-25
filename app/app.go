package app

import "fmt"

func Run() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6))
}

func Sum(nums ...int) int {
	i := 0
	for _, j := range nums {
		i += j
	}
	return i
}
