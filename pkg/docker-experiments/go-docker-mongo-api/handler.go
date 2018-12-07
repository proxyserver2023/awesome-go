package main

import (
	"fmt"
	"net/http"
)

func jobsGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string("Hello"))
}
