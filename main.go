package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"github.com/luho91/gato/internal/config"
	"github.com/luho91/gato/internal/database"
	"os"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Println("Error reading config file :(", err)
		return
	}

	s := state{}

	s.cfg = &cfg

	c := commands{}

	c.commands = make(map[string]func(*state, command) error)

	c.commands["login"] = handlerLogin
	
	db, err := sql.Open("postgres", s.cfg.DbURL)

	if err != nil {
		fmt.Printf("Something went wrong connecting to DB: %s", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	c.commands["register"] = handlerRegister

	s.db = dbQueries

	c.commands["reset"] = handlerReset
	c.commands["users"] = handlerGetUsers

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments passed.")
		os.Exit(1)
	}

	cName := args[1]
	cArgs := args[2:]

	cmd := command{}
	cmd.name = cName
	cmd.args = cArgs

	err = c.run(&s, cmd)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
