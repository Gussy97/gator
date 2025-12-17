package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	const url = "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(rssFeed)
	return nil
}
