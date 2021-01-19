package cmd

import "os"

const (
	cmdExit         = "exit"
	descriptionExit = "exit boltview"
)

type exit struct {
	base
}

func init() {
	RegisterCommand(cmdExit, &exit{})
}

func (e *exit) Open(opts ...ParseOption) (Command, error) {
	o := &exit{base: base{
		name:        cmdExit,
		cmd:         cmdExit,
		description: descriptionExit,
	}}
	for _, opt := range opts {
		opt(o)
	}
	return o, nil
}

func (e *exit) exec() error {
	writeToConsole("bye.\n")
	os.Exit(0)
	return nil
}
