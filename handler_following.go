package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	username := s.cfg.CurrentUserName

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), username)
	if err != nil {
		return err
	}

	fmt.Println("followed feeds:")
	for _, feed := range followedFeeds {
		fmt.Println(feed.Name)
	}

	return nil
}
