package cmd

import (
	"boltview/pkg/boltdb"
	"errors"
)

const (
	cmdKeys         = "keys"
	descriptionKeys = "show keys from the specific bucket"
)

func init() {
	RegisterCommand(cmdKeys, &key{})
}

type key struct {
	base
	bucket string
	filter []string
	keys   []string
}

func (k *key) Open(opts ...ParseOption) (Command, error) {
	o := &key{base: base{
		name:        cmdKeys,
		cmd:         cmdKeys,
		description: descriptionKeys,
	}}
	for _, opt := range opts {
		opt(o)
	}
	return o, nil
}

func (k *key) exec() error {
	var err error
	k.keys, err = boltdb.Keys(k.bucket)
	if err != nil {
		return err
	}
	return nil
}

func (k *key) parse(args []string) error {
	if len(args) < 2 {
		return errors.New("params is invalid")
	}
	k.bucket = args[1]
	k.filter = args[2:]
	return nil
}

func (k *key) ok() {
	writeToConsole(k.keys)
}
