package main

import (
	"github.com/Gussy97/gator/internal/config"
	"github.com/Gussy97/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
