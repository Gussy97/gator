package main

import (
	"context"
	"fmt"

	"github.com/Gussy97/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	user, err := s.db.GetUserByUsername(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	printFeedFollow(feed_follow)
	return nil
}

func printFeedFollow(feed_follow database.CreateFeedFollowRow) {
	fmt.Printf("Feed Name: %s\n", feed_follow.FeedName)
	fmt.Printf("User Name: %s\n", feed_follow.UserName)
}
