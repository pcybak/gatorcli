package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pcybak/gatorcli/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	userName := cmd.Arguments[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %s does not exist. register it first", userName)
	}
	if err != nil {
		return err
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

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("db error on GetUser: %v\n", err)
	}
	if err == nil {
		return fmt.Errorf("user already exists\n")
	}

	dateNow := time.Now()
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
		Name:      cmd.Arguments[0],
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

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)\n")
		} else {
			fmt.Println()
		}
	}
	return nil
}
