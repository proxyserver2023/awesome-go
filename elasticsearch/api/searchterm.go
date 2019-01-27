package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func searchWithATermQuery(client *elastic.Client, indexName string, ctx context.Context) *elastic.SearchResult {
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("tweets").            // search in index "tweets"
		Query(termQuery).           // specify the query
		Sort("user.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(ctx)                     // execute

	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	return searchResult
}
