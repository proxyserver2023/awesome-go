package main

import "context"

func createIndex(client *elastic.Client, indexName string, ctx context.Context) {
	_, err := client.CreateIndex(indexName).Do(ctx)
	if err != nil {

	}
}
