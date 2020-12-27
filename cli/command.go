package cli

import (
	"boltview/exec"
	"github.com/c-bata/go-prompt"
)

var cmdHistory []string

func Run() {
	for {
		t := prompt.Input("> ", completer, prompt.OptionHistory(cmdHistory))
		addHistory(t)
		exec.Run(t)
	}
}

func addHistory(s string) {
	cmdHistory = append(cmdHistory, s)
}
