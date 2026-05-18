package main

import (
	"github.com/luho91/gato/internal/config"
	"github.com/luho91/gato/internal/database"
)

type state struct {
	db		*database.Queries
	cfg		*config.Config
}
