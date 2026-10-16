package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pcybak/gatorcli/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}
	feedUrl := cmd.Arguments[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	dateNow := time.Now()
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("user: %s, feed: %s", feedFollow.Username, feedFollow.Feedname)

	return nil
}
