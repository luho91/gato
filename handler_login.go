package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/luho91/gato/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No username provided")
	}

	uName := cmd.args[0]
	s.cfg.CurrentUser = uName

	_, err := s.db.GetUser(context.Background(), uName)

	if err != nil {
		return err
	}

	err = s.cfg.SetUser(uName)

	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %s.\n", uName)

	return nil
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
	
}
