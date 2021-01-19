package cmd

import (
	"boltview/pkg/boltdb"
)

const (
	cmdCreate         = "create"
	descriptionCreate = "create buckets if does not exist"
)

type create struct {
	base
	bucketNames []string
}

func init() {
	RegisterCommand(cmdCreate, &create{})
}

func (c *create) Open(opts ...ParseOption) (Command, error) {
	o := &create{base: base{
		name:        cmdCreate,
		cmd:         cmdCreate,
		description: descriptionCreate,
	}}
	for _, opt := range opts {
		opt(o)
	}
	return o, nil
}

func (c *create) parse(args []string) error {
	if len(args) <= 1 {
		return nil
	}
	c.params = args
	c.bucketNames = c.params[1:]
	return nil
}

func (c *create) exec() error {
	for _, name := range c.bucketNames {
		err := boltdb.CreateBucket(name)
		if err != nil {
			return err
		}
	}
	return nil
}
