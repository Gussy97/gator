package main

import (
	"context"
	"fmt"

	"github.com/Gussy97/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUserByUsername(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting user: %w", err)
		}
		return handler(s, c, user)
	}
}
