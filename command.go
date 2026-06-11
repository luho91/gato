package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commands[cmd.name]

	if !ok {
		return errors.New("Command not found")
	}

	err := f(s, cmd)

	return err
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.commands[name] = f
}

func (c *commands) init() {
	c.commands = make(map[string]func(*state, command) error)
}

func (c *commands) registerAll() {
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerGetUsers)
	c.register("agg", middlewareLoggedIn(handlerAgg))
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", handlerFollow)
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
}
