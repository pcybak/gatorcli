package main

import (
	"context"
	"fmt"

	"github.com/pcybak/gatorcli/internal/database"
)

func handlerUnfollow(s *state, cmd command, u database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s, <url>", cmd.Name)
	}

	feedUrl := cmd.Arguments[0]
	feedFollowParams := database.DeleteFeedFollowParams{
		UserID: u.ID,
		Url:    feedUrl,
	}

	_, err := s.db.DeleteFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("unfollowed %s\n", feedUrl)

	return nil
}
