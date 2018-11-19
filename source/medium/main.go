package main

import "fmt"

func main() {
	{
		// Start Outer Block
		a := 1
		fmt.Println(a)
		{
			// Start Inner Block
			a := 2
			fmt.Println(a)
			// End   Inner Block
		}
		fmt.Println(a)
		// End   Outer Block
	}
}
