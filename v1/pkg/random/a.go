package random

import (
	"fmt"
	"math"
)

func arun() {
	const (
		Large = 3.725290298461914e-09
		Small = 1.0 / (1 << 28)
	)
	fmt.Println(Large, Small)
	if math.Abs(Large-Small) < 0.000000001 {
		fmt.Println("Wow! Did not expect that")
	}
}
