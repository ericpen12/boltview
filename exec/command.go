package exec

import (
	"boltview/boltdb"
	"fmt"
)

const (
	cmdCreate         = "create"
	descriptionCreate = "create buckets if does not exist"
)

type base struct {
	name        string
	cmd         string
	description string
	options     []string
	fn          func(name string) error
}

type create struct {
	base
	bucketNames []string
}

func (c create) Description() string {
	return c.description
}

func (c create) Parse(args []string) error {
	if len(c.options) <= 1 {
		return nil
	}
	c.options = args
	c.bucketNames = c.options[1:]
	return nil
}

func (c create) Exec() error {
	for _, name := range c.bucketNames {
		err := boltdb.CreateBucket(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c create) Error(err error) {
	fmt.Println(err)
}

func (c create) Ok() {
	printOK()
}

func (c create) CommandName() string {
	return c.name
}

func NewCmdCreate() *create {
	return &create{base{
		name:        cmdCreate,
		cmd:         cmdCreate,
		description: descriptionCreate,
		options:     nil,
	}, nil}
}

func printOK() {
	fmt.Println("ok")
}
