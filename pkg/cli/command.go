package cli

import (
	"boltview/cmd"
	c "boltview/pkg/cmd"
	"boltview/pkg/parser"
	"fmt"
	"github.com/c-bata/go-prompt"
)

var cmdHistory []string

func Run() {
	for {
		t := prompt.Input("> ", completer, prompt.OptionHistory(cmdHistory))
		addHistory(t)
		if cmd.Release {
			p, err := c.NewParser(t)
			if err != nil {
				fmt.Println(err)
				continue
			}
			p.Parse(t)
		} else {
			parser.Run(t)
		}
	}
}

func addHistory(s string) {
	cmdHistory = append(cmdHistory, s)
}
