package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	rssFeedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), rssFeedURL)
	if err != nil {
		return fmt.Errorf("error while fetching RSS feed from %s: %v", rssFeedURL, err)
	}

	prettyjson, _ := json.MarshalIndent(*feed, "", "  ")
	fmt.Println(string(prettyjson))

	return nil
}
