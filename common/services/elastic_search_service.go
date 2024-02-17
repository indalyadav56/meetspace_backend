package services

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

func ElasticSearchService() {

    // Create Elasticsearch client
    client, err := elastic.NewClient(
        elastic.SetURL("http://localhost:9200"), 
        elastic.SetSniff(false))
    if err != nil {
        panic(err)
    }

    // Create a document
    doc := elastic.NewBulkIndexRequest().Index("users").Id("1").Doc(map[string]interface{}{
        "name":    "John",
        "gender":  "male",
        "age":     30,
    })
	
    // Add document to Elasticsearch
    _, err = client.Bulk().Add(doc).Do(context.Background())
    if err != nil {
        panic(err)
    }
	
    // Search for document
    termQuery := elastic.NewTermQuery("name.keyword", "John")
    searchResult, err := client.Search().
        Index("users").
        Query(termQuery).
        Pretty(true).
        Do(context.Background())
    if err != nil {
        panic(err) 
    }

    fmt.Printf("Found %d user with name John\n", searchResult.TotalHits())

    // Close Elasticsearch connection
    client.Stop()	
}
