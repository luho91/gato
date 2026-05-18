package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/luho91/gato/internal/database"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No name provided")
	}

	name := cmd.args[0]

	p := database.CreateUserParams{}
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Name = name

	_, err := s.db.CreateUser(context.Background(), p)

	if err != nil {
		return err
	}

	s.cfg.CurrentUser = name
	err = s.cfg.SetUser(name)

	if err != nil {
		return err
	}

	fmt.Printf("User has been created: %s\n", name)

	return nil
}
