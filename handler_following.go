package main

import (
	"context"
	"fmt"

	"github.com/pcybak/gatorcli/internal/database"
)

func handlerFollowing(s *state, cmd command, u database.User) error {
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), u.Name)
	if err != nil {
		return err
	}

	fmt.Println("followed feeds:")
	for _, feed := range followedFeeds {
		fmt.Println(feed.Name)
	}

	return nil
}
