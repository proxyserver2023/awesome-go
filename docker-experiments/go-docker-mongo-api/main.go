package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. get config from env
	// 2. register Database
	// 3. register the routes
	router := NewRouter()
	// 4. listen and serve http server
	log.Fatal(http.ListenAndServe(":9090", router))
}
