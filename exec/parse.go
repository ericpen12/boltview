package exec

import "errors"

var (
	ErrCommandExist = errors.New("command already exist")
)

var commandMap map[string]Command

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

type Command interface {
	CommandName() string
	Description() string
	Error(err error)
	Exec() error
	Parse() error
	Ok()
}
