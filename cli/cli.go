package cli

import (
	"github.com/c-bata/go-prompt"
)

var s = []prompt.Suggest{
	{Text: "buckets", Description: "List all buckets"},
	{Text: "keys", Description: "Get all keys according to buckets"},
	{Text: "get", Description: "show value"},
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
