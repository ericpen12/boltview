package cli

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

var remindMap = map[string][]prompt.Suggest{
	cmdGet: {
		{Text: "all"},
		{Text: "*"},
	},
}

func addSuggests(key string, val []string) {
	var su []prompt.Suggest
	for _, v := range val {
		su = append(su, prompt.Suggest{Text: v})
	}
	remindMap[key] = append(remindMap[key], su...)
}

var defaultSugget = []prompt.Suggest{
	{Text: "buckets", Description: "List all buckets"},
	{Text: "keys", Description: "Get all keys according to buckets"},
	{Text: "get", Description: "show value"},
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(currentRemind(d), d.GetWordBeforeCursor(), true)
}

func currentRemind(d prompt.Document) []prompt.Suggest {
	c := strings.Trim(d.Text, " ")
	if _, ok := remindMap[c]; ok || len(c) < len(d.Text) {
		//log.Print(val)
		return nil
	}
	return defaultSugget
}
