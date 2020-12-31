package cli

import (
	"boltview/boltdb"
	"fmt"
	"github.com/c-bata/go-prompt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	commandNotFound = "command not found:"

	cmdBuckets = "buckets"
	cmdKeys    = "keys"
	cmdGet     = "get"
	cmdQ       = "q"
	cmdExit    = "exit"
	cmdBye     = "bye"
)

var mapFunc = map[string]func(c *cmd) error{
	cmdBuckets: buckets,
	cmdKeys:    keys,
	cmdGet:     get,
	cmdQ:       exit,
	cmdExit:    exit,
	cmdBye:     exit,
}

var cmdHistory []string

func Run() {
	for {
		t := prompt.Input("> ", completer, prompt.OptionHistory(cmdHistory))
		cmd := parseCmd(t)
		if fn, ok := mapFunc[cmd.fn]; ok {
			addHistory(cmd)
			fn(cmd)
		} else {
			fmt.Println(commandNotFound, t)
		}
	}
}

func addHistory(cmd *cmd) {
	cmdHistory = append(cmdHistory, cmd.fn+" "+strings.Join(cmd.options, " "))
}

func buckets(c *cmd) error {
	bu, err := boltdb.Buckets()
	if err != nil {
		return err
	}
	addSuggests(cmdKeys, bu)
	fmt.Println(bu)
	return nil
}

func keys(c *cmd) error {
	data, err := boltdb.Keys(c.options[0])
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func get(c *cmd) error {
	data, err := boltdb.Get(c.options[0])
	if err != nil {
		return err
	}
	if len(c.options) > 2 {
		if c.options[1] == "-f" {
			err := ioutil.WriteFile(c.options[2], data, 0700)
			if err != nil {
				return err
			}
			fmt.Println("ok")
			return nil
		}
	}
	fmt.Println(string(data))
	return nil
}

type cmd struct {
	fn      string
	options []string
}

func parseCmd(input string) *cmd {
	c := new(cmd)
	s := strings.Split(input, " ")

	if len(s) > 0 {
		c.fn = s[0]
	}
	c.options = s[1:]
	return c
}

func exit(*cmd) error {
	os.Exit(0)
	return nil
}
