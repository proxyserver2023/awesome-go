package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func getESVersion(client *elastic.Client, ctx context.Context) string {
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	return fmt.Sprintf("Elasticsearch version %s\n", esversion)
}
