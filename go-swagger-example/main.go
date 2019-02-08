package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
}
