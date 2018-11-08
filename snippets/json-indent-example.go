package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}

	Roads := []Road{
		{"Alamin", 1},
		{"BCD", 2},
	}

	b, err := json.Marshal(Roads)

	if err != nil {
		log.Fatal("Could Not Serialize to JSON")
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")

	out.WriteTo(os.Stdout)
}
