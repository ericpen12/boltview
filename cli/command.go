package cli

import (
	"boltview/boltdb"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapFunc = map[string]func(c *cmd) error{
	"bucket": buckets,
	"keys":   keys,
	"get":    get,
}

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		cmd := parseCmd(scanner.Text())
		if fn, ok := mapFunc[cmd.fn]; ok {
			fn(cmd)
		} else {
			fmt.Println(scanner.Text(), "not exist")
		}
		if scanner.Text() == "q" {
			break
		}
		fmt.Print(">>> ")
	}
}

func buckets(c *cmd) error {
	bu, err := boltdb.Buckets()
	if err != nil {
		return err
	}
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
