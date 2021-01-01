package cli

import (
	"boltview/pkg/parser"
	"github.com/c-bata/go-prompt"
	"strings"
)

const (
	cmdMode = 1 + iota
	optionMode
)

var defaultSuggest []prompt.Suggest

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(
		suggest(d),
		d.GetWordBeforeCursor(),
		true,
	)
}

func suggest(d prompt.Document) []prompt.Suggest {
	if d.Text == "" {
		return nil
	}
	mode := len(strings.Split(d.Text, " "))
	switch mode {
	case cmdMode:
		return commandSuggest()
	}
	return nil
}

func commandSuggest() []prompt.Suggest {
	if defaultSuggest == nil {
		list := parser.CommandList()
		defaultSuggest = make([]prompt.Suggest, len(list))
		for i, c := range list {
			defaultSuggest[i] = prompt.Suggest{
				Text:        c.CommandName(),
				Description: c.Description(),
			}
		}
	}
	return defaultSuggest
}
