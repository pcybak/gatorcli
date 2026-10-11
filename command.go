package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pcybak/gatorcli/internal/database"
)

type command struct {
	Name      string
	Arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	userName := cmd.Arguments[0]
	_, err := s.db.GetUser(context.Background(), sql.NullString{String: userName, Valid: true})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Fatalf("user %s does not exist. register it first", userName)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to set username to provided one: %v\n", err)
	}
	fmt.Printf("set username to: %v\n", userName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	userName := cmd.Arguments[0]
	sqlUserName := sql.NullString{
		String: userName,
		Valid:  true,
	}

	_, err := s.db.GetUser(context.Background(), sqlUserName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalf("db error on GetUser: %v\n", err)
	}
	if err == nil {
		log.Fatalf("user already exists\n")
	}

	dateNow := time.Now()
	userParams := database.CreateUserParams{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		CreatedAt: sql.NullTime{
			Time:  dateNow,
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  dateNow,
			Valid: true,
		},
		Name: sql.NullString{
			String: cmd.Arguments[0],
			Valid:  true,
		},
	}
	_, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	loginCmd := command{
		Name:      "login",
		Arguments: cmd.Arguments,
	}
	err = handlerLogin(s, loginCmd)
	if err != nil {
		return err
	}

	fmt.Printf("created user %v\n", userName)

	return nil
}

type commands struct {
	cmd map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if val, ok := c.cmd[cmd.Name]; ok {
		val(s, cmd)
		return nil
	}
	return fmt.Errorf("no such function: %v", cmd.Name)
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.cmd[name] = f
}
