package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ioUtil() {
	r := strings.NewReader(
		"Go is a general purpose",
		)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}

func ioUtilReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	content, err := ioutil.ReadFile("Makefile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File Contents: %s", content)
}

