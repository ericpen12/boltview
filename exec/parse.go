package exec

import (
	"errors"
	"fmt"
	"strings"
)

const (
	commandNotFound = "command not found:"

	cmdCreate         = "create"
	descriptionCreate = "create buckets if does not exist"
)

var (
	ErrCommandExist = errors.New("command already exist")
)

type base struct {
	name        string
	cmd         string
	description string
	options     []string
	fn          func(name string) error
}

type Command interface {
	CommandName() string
	Description() string
	Error(err error)
	Exec() error
	Parse(args []string) error
	Ok()
}

var commandMap = map[string]Command{}

func init() {
	addCmd(NewCmdCreate())
}

func addCmd(c Command) error {
	if _, ok := commandMap[c.CommandName()]; ok {
		return ErrCommandExist
	}

	commandMap[c.CommandName()] = c
	return nil
}

func Run(s string) {
	args := strings.Split(s, " ")
	if len(args) < 1 {
		return
	}
	c, ok := commandMap[args[0]]
	if !ok {
		fmt.Println(commandNotFound, args[0])
		return
	}
	if err := c.Parse(args); err != nil {
		c.Error(err)
	}

	if err := c.Exec(); err != nil {
		c.Error(err)
		return
	}
	c.Ok()
}

func print(s string) {
	fmt.Println(s)
}

func printOK() {
	print("ok")
}
