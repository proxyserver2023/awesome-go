package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func pingClient(client *elastic.Client, ctx context.Context) {
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle Error
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
