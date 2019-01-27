package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

func main() {
	//router := mux.NewRouter()
	//router.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Hello There")
	//})
	//http.Handle("/", router)
	//http.ListenAndServe(":8081", nil)

	client := connectElasticSearch()
	pingClient(client, context.Background())
	fmt.Println("ES Version =>", getESVersion(client, context.Background()))

	if exists := checkIndex(client, "tweets", context.Background()); !exists {
		fmt.Println("Exists")
	} else {
		createIndex(client, "tweets", context.Background())
	}

	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	indexTweet(client, "tweets", context.Background(), "1", &tweet1)

	tweet2 := Tweet{User: "alamin", Message: "Six Days", Retweets: 1}
	indexTweet(client, "tweets", context.Background(), "2", &tweet2)

	searchResult := searchWithATermQuery(client, "tweets", context.Background())
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		fmt.Print("Found no tweets\n")
	}

	// Delete the index again
	_, err := client.DeleteIndex("tweets").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

}
