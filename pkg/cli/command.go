package cli

import (
	"boltview/pkg/parser"
	"github.com/c-bata/go-prompt"
)

var cmdHistory []string

func Run() {
	for {
		t := prompt.Input("> ", completer, prompt.OptionHistory(cmdHistory))
		addHistory(t)
		parser.Run(t)
	}
}

func addHistory(s string) {
	cmdHistory = append(cmdHistory, s)
}
