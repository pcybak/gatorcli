package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pcybak/gatorcli/internal/config"
	"github.com/pcybak/gatorcli/internal/database"
)

var (
	programState state
	c            commands
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not load config file: %v\n", err)
		os.Exit(1)
	}

	programState.cfg = &cfg
	db, err := sql.Open("postgres", programState.cfg.DbUrl)
	if err != nil {
		log.Fatalf("cannot connect to database %v: %v", programState.cfg.DbUrl, err)
	}
	dbQueries := database.New(db)
	programState.db = dbQueries
	c.cmd = make(map[string]func(*state, command) error)

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("browse", middlewareLoggedIn(handleBrowse))

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("no function name was provided\n")
	}
	commandName := args[1]
	cmd := command{
		Name:      commandName,
		Arguments: os.Args[2:],
	}
	err = c.run(&programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
