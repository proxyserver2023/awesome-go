package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func indexTweet(client *elastic.Client, indexName string, ctx context.Context, id string, tweet *Tweet) {
	put1, err := client.Index().
		Index(indexName).
		Type("doc").
		Id(id).
		BodyJson(*tweet).
		Refresh("wait_for").
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
