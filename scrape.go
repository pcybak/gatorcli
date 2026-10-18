package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pcybak/gatorcli/internal/database"
)

func scrapeFeeds(s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeedToFetch.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	rssFeed, err := fetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		return err
	}

	fmt.Println("Fetched topics:")
	for _, feed := range rssFeed.Channel.Item {
		fmt.Println(feed.Title)
	}

	return nil
}
