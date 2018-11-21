package golang

import "net/http"

func StripPrefix() {
	http.Handle(
		"/tmpfiles",
		http.StripPrefix("/tmpfiles", http.FileServer(http.Dir("/tmp"))),
	)
}
