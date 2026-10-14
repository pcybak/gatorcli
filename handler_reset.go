package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUser(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("deleted all records from table users\n")
	return nil
}
