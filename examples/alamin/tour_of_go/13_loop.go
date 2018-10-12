package main

import (
	"fmt"
	"math"
)

const eps = 0.0000000000000000000000000001

func sqrt(x float64) float64 {
	z := 1.0
	i := 1
	for ; math.Abs(math.Sqrt(x)-z) > eps; z, i = z-(z*z-x)/(2*z), i+1 {
		fmt.Println(i, "=>", z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}
