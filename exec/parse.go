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
	error(err error)
	exec() error
	parse(args []string) error
	ok()
}

var commandMap = map[string]Command{}

func CommandList() []Command {
	result := make([]Command, len(commandMap))
	var i int
	for _, c := range commandMap {
		result[i] = c
		i++
	}
	return result
}

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
	params      []string
	fn          func(name string) error
}

func (b *base) Exec() error {
	return nil
}

func (b *base) parse([]string) error {
	return nil
}

func (b *base) CommandName() string {
	return b.name
}

func (b *base) Description() string {
	return b.description
}

func (b *base) ok() {
	writeToConsole("ok")
}

func (b *base) error(err error) {
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
	if err := c.parse(args); err != nil {
		c.error(err)
		return
	}

	if err := c.exec(); err != nil {
		c.error(err)
		return
	}
	c.ok()
}

func writeToConsole(s ...interface{}) {
	fmt.Println(s...)
}
