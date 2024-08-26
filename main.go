package main

import (
	"fmt"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.New("http://localhost:7700")
	index := client.Index("movies")
	err := addDocuments(index)
	if err != nil {
		fmt.Println("add documents failed", err)
	}

	// 更新索引设置
	_, err = index.UpdateSettings(&meilisearch.Settings{
		FilterableAttributes: []string{"genres"},
	})
	if err != nil {
		fmt.Println("update settings failed", err)
	}

	time.Sleep(time.Second * 2)
	searchRes, err := custom_search(index)
	if err != nil {
		fmt.Println("search failed", err)
	}
	dump.P(searchRes)
}

func addDocuments(index meilisearch.IndexManager) error {
	documents := []map[string]interface{}{
		{"id": 1, "title": "Carol01", "genres": []string{"Romance", "Drama"}},
		{"id": 2, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 3, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 4, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 5, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 6, "title": "Philadelphia", "genres": []string{"Drama"}},
	}

	_, err := index.AddDocuments(documents)
	if err != nil {
		return err
	}
	return nil
}

func search(index meilisearch.IndexManager) (*meilisearch.SearchResponse, error) {
	return index.Search("philoudelphia",
		&meilisearch.SearchRequest{
			Limit: 10,
		})
}

func custom_search(index meilisearch.IndexManager) (*meilisearch.SearchResponse, error) {
	return index.Search("",
		&meilisearch.SearchRequest{
			Filter: []string{"genres = action"},
		})
}
