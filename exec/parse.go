package exec

import (
	"fmt"
	"strings"
)

const (
	commandNotFound = "command not found:"
)

type Command interface {
	CommandName() string
	Description() string
	Error(err error)
	Exec() error
	Parse(args []string) error
	Ok()
}

var commandMap = map[string]Command{}

func register(c Command) {
	_, ok := commandMap[c.CommandName()]
	if !ok {
		commandMap[c.CommandName()] = c
	}
}

type base struct {
	name        string
	cmd         string
	description string
	options     []string
	fn          func(name string) error
}

func (b *base) Exec() error {
	return nil
}

func (b *base) Parse([]string) error {
	return nil
}

func (b *base) CommandName() string {
	return b.name
}

func (b *base) Description() string {
	return b.description
}

func (b *base) Ok() {
	writeToConsole("ok")
}

func (b *base) Error(err error) {
	writeToConsole(err)
}

func Run(s string) {
	args := strings.Split(s, " ")
	if len(args) < 1 {
		return
	}
	c, ok := commandMap[args[0]]
	if !ok {
		writeToConsole(commandNotFound, args[0])
		return
	}
	if err := c.Parse(args); err != nil {
		c.Error(err)
		return
	}

	if err := c.Exec(); err != nil {
		c.Error(err)
		return
	}
	c.Ok()
}

func writeToConsole(s ...interface{}) {
	fmt.Println(s...)
}
