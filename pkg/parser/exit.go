package parser

import "os"

const (
	cmdExit         = "exit"
	descriptionExit = "exit boltview"
)

type exit struct {
	base
}

func init() {
	register(newExit())
}

func newExit() *exit {
	return &exit{
		base{
			name:        cmdExit,
			cmd:         cmdExit,
			description: descriptionExit,
		},
	}
}

func (e *exit) exec() error {
	writeToConsole("bye.\n")
	os.Exit(0)
	return nil
}
