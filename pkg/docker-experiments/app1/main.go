package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // No Password
		DB:       0,  // Use Default DB
	})

	pong, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(pong)

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// http.Handle("/", r)
	// http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
