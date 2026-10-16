package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	prettyjson, _ := json.MarshalIndent(feeds, "", "  ")
	fmt.Println(string(prettyjson))

	return nil
}
