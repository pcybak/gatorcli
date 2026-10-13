package main

import (
	"fmt"
)

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if val, ok := c.cmd[cmd.Name]; ok {
		return val(s, cmd)
	}
	return fmt.Errorf("no such function: %v", cmd.Name)
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.cmd[name] = f
}
