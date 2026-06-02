package main

import (
	"context"
	"errors"
	"github.com/luho91/gato/internal/database"
	"github.com/google/uuid"
	"time"
	"fmt"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No URL provided")
	}

	url := cmd.args[0]

	r, err := subscribeFeed(s, url)
	if err != nil {
		return err
	}

	fmt.Printf("Username: %v\n", r.UserName)
	fmt.Printf("Feedname: %v\n", r.FeedName)

	return nil
}

func subscribeFeed(s *state, feedURL string) (database.CreateFeedFollowRow, error) {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return database.CreateFeedFollowRow{}, err
	}

	feed, err := s.db.GetOneFeedByURL(context.Background(), feedURL)
	if err != nil {
		return database.CreateFeedFollowRow{}, err
	}

	p := database.CreateFeedFollowParams{}
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.UserID = user.ID
	p.FeedID = feed.ID

	r, err := s.db.CreateFeedFollow(context.Background(), p)

	return r, err
	
}
