package main

import "context"
import (
	"github.com/olivere/elastic"
)

func createIndex(client *elastic.Client, indexName string, ctx context.Context) {
	_, err := client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		// Handle ERrrror
	}
}
