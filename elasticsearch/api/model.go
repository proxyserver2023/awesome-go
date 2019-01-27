package main

import (
	"time"

	"github.com/olivere/elastic"
)

//type Document struct {
//	ID        string    `json:"id"`
//	Title     string    `json:"title"`
//	CreatedAt time.Time `json:"created_at"`
//	Content   string    `json:"content"`
//}
//
//type DocumentRequest struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//}
//
//type DocumentResponse struct {
//	Title     string    `json:"title"`
//	CreatedAt time.Time `json:"created_at"`
//	Content   string    `json:"content"`
//}
//
//type SearchResponse struct {
//	Time      string             `json:"time"`
//	Hits      string             `json:"hits"`
//	Documents []DocumentResponse `json:"documents"`
//}

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
   "settings":{
      "number_of_shards":1,
      "number_of_replicas":0
   },
   "mappings":{
      "tweet":{
         "properties":{
            "user":{
               "type":"keyword"
            },
            "message":{
               "type":"text",
               "store":true,
               "fielddata":true
            },
            "image":{
               "type":"keyword"
            },
            "created":{
               "type":"date"
            },
            "tags":{
               "type":"keyword"
            },
            "location":{
               "type":"geo_point"
            },
            "suggest_field":{
               "type":"completion"
            }
         }
      }
   }
}

`
