package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Started")
	defer log.Printf("Handler ended")

	// time.Sleep(5*time.Second)
	fmt.Fprintln(w, "hello")
}
