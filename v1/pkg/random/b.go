package random

import (
	"fmt"
)

func brun() {

	switch shit := "pdf"; shit {
	case "pdf":
		fmt.Println("PDF found")
	case "csv":
		fmt.Println("CSV found")
	case "text":
		fmt.Println("text found")
	default:
		fmt.Println("Nothing Found")
	}
}
