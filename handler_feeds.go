package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, f := range feeds {
		u, err := s.db.GetUserByUUID(context.Background(), f.UserID)

		if err != nil {
			return err
		}
		fmt.Printf("%v %v %v\n", f.Name, f.Url, u.Name)
	}

	return nil
}
