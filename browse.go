package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pcybak/gatorcli/internal/database"
)

func handleBrowse(s *state, cmd command, u database.User) error {
	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	var err error
	limit := 2

	if len(cmd.Arguments) == 1 {
		limit, err = strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return fmt.Errorf("limit has to be a number")
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: u.ID, Limit: int32(limit)})
	if err != nil {
		return err
	}
	fmt.Println("Your posts:")
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.PublishedAt)
		fmt.Println(post.Description)
		fmt.Println(post.Url)
	}

	return nil
}
