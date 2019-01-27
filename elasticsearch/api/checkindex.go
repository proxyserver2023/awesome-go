package main

import (
	"context"

	"github.com/olivere/elastic"
)

func checkIndex(client *elastic.Client, indexName string, ctx context.Context) bool {
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		// Handle Error
	}
	return exists
}
