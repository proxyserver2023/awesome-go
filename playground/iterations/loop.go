package iterations

import "fmt"

var hello = []int{1,2,3,4,5,6}

func Run() {
	for _, i := range hello {
		fmt.Println(i)
	}
}