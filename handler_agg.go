package main

import (
	"context"
	"fmt"
	"github.com/luho91/gato/internal/database"
	"github.com/google/uuid"
	"time"
	"errors"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("No interval provided")
	}

	t, err := time.ParseDuration(cmd.args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", t)

	ticker := time.NewTicker(t)
	for ;; <-ticker.C {
		err := scrapeFeeds(s, user.ID)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state, userID uuid.UUID) error {
	f, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	p1 := database.MarkFeedFetchedParams{}
	p1.LastFetchedAt.Time = time.Now()
	p1.LastFetchedAt.Valid = true
	p1.ID = f.ID
	s.db.MarkFeedFetched(context.Background(), p1)
	feed, err := fetchFeed(context.Background(), f.Url)
	if err != nil {
		return err
	}

	for _, i := range feed.Channel.Item {
		p2 := database.CreatePostParams{}
		p2.ID = uuid.New()
		p2.CreatedAt = time.Now()
		p2.UpdatedAt = time.Now()
		p2.Title = i.Title
		p2.Url = f.Url
		p2.FeedID = f.ID

		p, _ := s.db.CreatePost(context.Background(), p2)

		fmt.Printf("%v\n", p.Title)
	}

	fmt.Printf("Fetched %v\n", f.Url)

	return nil
}
