package exec

import (
	"boltview/boltdb"
	"errors"
)

const (
	cmdKeys         = "keys"
	descriptionKeys = ""
)

type key struct {
	base
	bucket string
	filter []string
	keys   []string
}

func newKey() *key {
	return &key{base: base{
		name:        cmdKeys,
		cmd:         cmdKeys,
		description: descriptionKeys,
	}}
}

func (k *key) CommandName() string {
	return k.name
}

func (k *key) Description() string {
	return k.description
}

func (k *key) Error(err error) {
}

func (k *key) Exec() error {
	var err error
	k.keys, err = boltdb.Keys(k.bucket)
	if err != nil {
		return err
	}
	return nil
}

func (k *key) Parse(args []string) error {
	if len(args) < 2 {
		return errors.New("params is invalid")
	}
	k.bucket = args[1]
	k.filter = args[2:]
	return nil
}

func (k *key) Ok() {
	print(k.keys)
}