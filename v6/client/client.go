package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := "http://localhost:8080"
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
		)

	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
}
