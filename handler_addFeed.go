package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pcybak/gatorcli/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feedName := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	dateNow := time.Now()
	feedRecord := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := s.db.AddFeed(context.Background(), feedRecord)
	if err != nil {
		return err
	}

	prettyjson, _ := json.MarshalIndent(feed, "", "  ")
	fmt.Println(string(prettyjson))

	return nil
}
