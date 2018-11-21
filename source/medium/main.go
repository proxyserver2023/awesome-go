package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func sleep() {
	time.Sleep(
		time.Duration(rand.Intn(1000)) * time.Millisecond,
	)
}

func reader() {

}

func writer() {

}

func main() {

}
