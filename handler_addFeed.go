package main

import (
	"context"
	"fmt"

	"github.com/Gussy97/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	dbUser, err := s.db.GetUserByUsername(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: dbUser.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
