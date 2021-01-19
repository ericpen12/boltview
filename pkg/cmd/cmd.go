package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

const (
	commandNotFound  = "command not found: %s"
	commandSeparator = " "
)

type ErrCmdNotFound struct {
	Msg string
}

func (e *ErrCmdNotFound) Error() string {
	return fmt.Sprintf(commandNotFound, e.Msg)
}

type CommandDriver interface {
	Open(opts ...ParseOption) (Command, error)
}

type Command interface {
	CommandName() string
	Description() string
	error(err error)
	exec() error
	parse(args []string) error
	ok()
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

var (
	cmdMu sync.RWMutex
	cmd   = make(map[string]CommandDriver)
)

func NewCommand(input string, opts ...ParseOption) (Command, error) {
	if len(input) == 0 {
		return nil, errors.New("empty input")
	}
	cmdStr := strings.Split(input, commandSeparator)[0]
	cmdMu.RLock()
	defer cmdMu.RUnlock()
	c, ok := cmd[cmdStr]
	if !ok {
		return nil, &ErrCmdNotFound{Msg: cmdStr}
	}
	return c.Open(opts...)
}

func writeToConsole(s ...interface{}) {
	fmt.Println(s...)
}
