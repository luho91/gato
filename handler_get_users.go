package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return err
	}

	for _, u := range users {
		out := fmt.Sprintf("* %s", u.Name)

		if u.Name == s.cfg.CurrentUser {
			out += " (current)"
		}

		fmt.Println(out)
	}

	return nil
}
