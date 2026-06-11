package main

import (
	"context"
	"github.com/luho91/gato/internal/database"
	"errors"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("No URL provided")
	}

	url := cmd.args[0]

	feed, err := s.db.GetOneFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	d := database.DeleteFollowParams{}

	d.UserID = user.ID
	d.FeedID = feed.ID

	err = s.db.DeleteFollow(context.Background(), d)

	return err
}
