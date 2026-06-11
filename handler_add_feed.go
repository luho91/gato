package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/luho91/gato/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("2 args required: name and url (separated by space)")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	p := database.CreateFeedParams{}
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Name = name
	p.Url = url
	p.UserID = user.ID

	_, err := s.db.CreateFeed(context.Background(), p)
	if err != nil {
		return err
	}

	_, err = subscribeFeed(s, url)
	
	return err
}
