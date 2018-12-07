package simpleapi

import (
	"fmt"
	"net/http"
)

// auth
func Login(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Hello There "+vars["user"])
}

// www
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am in HomePage")
}

// order

// product
