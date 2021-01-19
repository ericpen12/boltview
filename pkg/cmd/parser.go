package cmd

import (
	"strings"
)

type Parser interface {
	Parse(xml string) error
	Commands() []byte
	ReadChan() <-chan interface{}
}

type CommandParser struct {
	cmd    Command
	outPut string
}
type ParseOption func(opts interface{})

func NewParser(input string, opts ...ParseOption) (Parser, error) {
	c, err := NewCommand(input, opts...)
	if err != nil {
		return nil, err
	}
	return &CommandParser{cmd: c}, nil
}

func (c *CommandParser) Parse(input string) error {
	if err := c.cmd.parse(strings.Split(input, commandSeparator)); err != nil {
		return err
	}
	if err := c.cmd.exec(); err != nil {
		return err
	}
	c.cmd.ok()
	return nil
}

func (c *CommandParser) Commands() []byte {
	panic("implement me")
}

func (c *CommandParser) ReadChan() <-chan interface{} {
	panic("implement me")
}

func RegisterCommand(name string, driver CommandDriver) {
	cmdMu.Lock()
	defer cmdMu.Unlock()
	if driver == nil {
		panic("raster: Register Raster Driver is nil")
	}
	if _, ok := cmd[name]; ok {
		panic("command already register: " + name)
	}
	cmd[name] = driver
}
