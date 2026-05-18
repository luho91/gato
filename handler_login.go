package main

import (
	"context"
	"errors"
	"fmt"
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
