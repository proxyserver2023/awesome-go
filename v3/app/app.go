package app

import "fmt"

// Run - Driver
func Run() {
	if err := LoadConfig("v3/config"); err != nil {
		panic(fmt.Errorf("Invalid Application Configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages
}
