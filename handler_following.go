package main

import (
	"context"
	"fmt"
	"github.com/luho91/gato/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range feeds {
		fmt.Printf("%v\n", f.FeedName)
	}

	return nil
}
