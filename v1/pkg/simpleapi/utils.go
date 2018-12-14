package simpleapi

import "fmt"

func CheckErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
