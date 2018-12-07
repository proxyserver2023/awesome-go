package gopl

import "fmt"

func CheckErrPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
