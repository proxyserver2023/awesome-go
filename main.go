package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{
		"alamin",
		"I want something like this.",
		time.Now().Unix(),
	}

	byt, _ := json.Marshal(m)
	fmt.Println(byt)

	n := Message{}
	_ = json.Unmarshal(byt, &n)
	fmt.Println(n)
}
