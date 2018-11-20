package main

import (
	"log"
	"net/http"
)

func FileServer() {
	log.Fatal(
		http.ListenAndServe(
			":8080",
			http.FileServer(
				http.Dir("/usr/share/doc"),
			),
		),
	)
}
