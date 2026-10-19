package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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

	for _, feed := range rssFeed.Channel.Item {
		dateNow := time.Now()

		pubDate, err := time.Parse(time.RFC822, feed.PubDate)
		if err != nil {
			pubDate, err = time.Parse(time.RFC1123Z, feed.PubDate)
		}
		publishedAt := sql.NullTime{
			Time:  pubDate,
			Valid: err == nil,
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   dateNow,
			UpdatedAt:   dateNow,
			Title:       feed.Title,
			Url:         feed.Link,
			Description: feed.Description,
			PublishedAt: publishedAt,
			FeedID:      nextFeedToFetch.ID,
		})
		if err != nil {
			fmt.Printf("error during scraping feed %s: %v", nextFeedToFetch.Url, err)
		}
	}

	return nil
}
