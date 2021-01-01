package parser

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
	register(newCreate())
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

func newCreate() *create {
	return &create{base: base{
		name:        cmdCreate,
		cmd:         cmdCreate,
		description: descriptionCreate,
	}}
}
