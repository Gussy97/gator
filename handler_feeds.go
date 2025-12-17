package main

import (
	"context"
	"fmt"

	"github.com/Gussy97/gator/internal/database"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		printFeed(feed, user.Name)
	}
	return nil
}

func printFeed(feed database.Feed, feedUsername string) {
	fmt.Printf("ID:       %v\n", feed.ID)
	fmt.Printf("Name:     %s\n", feed.Name)
	fmt.Printf("URL:      %s\n", feed.Url)
	fmt.Printf("Added by: %s\n", feedUsername)
}
