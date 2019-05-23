package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

func spawnAWebServer() {
	router := mux.NewRouter()
	router.HandleFunc("/tasty", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Println("Got Response from server")
	})
	go http.ListenAndServe(":18000", router)
}

func main() {
	spawnAWebServer()
	time.Sleep(1 * time.Second)
	request, reqErr := http.NewRequest("GET", "http://localhost:181000/tasty", nil)
	if reqErr != nil {
		fmt.Println("Hello There")
		return
	}
	client := &http.Client{
		Timeout: time.Duration(2 * time.Second),
	}

	response, err := client.Do(request)
	if err != nil {
		switch err := err.(type) {
		case net.Error:
			if err.Timeout() {
				fmt.Println("This was a net.Error with a Timeout")
			}
		case *url.Error:
			fmt.Println("This is a *url.Error")
			if err, ok := err.Err.(net.Error); ok && err.Timeout() {
				fmt.Println("and it was because of a timeout")
			}
		}
		fmt.Println(err)
		return
	}

	if response.StatusCode == http.StatusOK {
		fmt.Println("Doing OK")
	} else {
		fmt.Println("Without OK - %d", response.StatusCode)
	}
}

// func sendJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
