package main

import (
	"github.com/olivere/elastic"
)

func connectElasticSearch() *elastic.Client {
	client, err := elastic.NewClient()
	if err != nil {
		// Handle Error
	}
	return client
}
