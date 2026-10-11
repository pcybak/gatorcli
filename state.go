package main

import (
	"github.com/pcybak/gatorcli/internal/config"
	"github.com/pcybak/gatorcli/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
