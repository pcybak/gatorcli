package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <duration>", cmd.Name)
	}

	refreshDurationString := cmd.Arguments[0]
	refreshDuration, err := time.ParseDuration(refreshDurationString)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(refreshDuration)
	defer ticker.Stop()

	fmt.Printf("Collecting feeds every %s\n", refreshDuration.String())

	scrapeFeeds(s)
	for {
		select {
		case <-ticker.C:
			scrapeFeeds(s)
		case <-context.Background().Done():
			return nil
		}
	}
}
