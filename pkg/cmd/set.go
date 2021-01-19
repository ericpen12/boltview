package cmd

import (
	"boltview/pkg/boltdb"
	"errors"
)

const (
	cmdSet         = "set"
	descriptionSet = "create a specific k/v"
)

var (
	ErrInvalidParams = errors.New("invalid params")
)

func init() {
	RegisterCommand(cmdSet, &set{})
}

type set struct {
	base
	bucket string
	val    map[string]string
}

func (s *set) Open(opts ...ParseOption) (Command, error) {
	o := &set{base: base{
		name:        cmdSet,
		cmd:         cmdSet,
		description: descriptionSet,
	}}
	for _, opt := range opts {
		opt(o)
	}
	return o, nil
}

func (s *set) exec() error {
	for k, v := range s.val {
		err := boltdb.Set(s.bucket, k, []byte(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *set) parse(args []string) error {
	num := len(args)
	if num%2 != 0 || num < 4 {
		return ErrInvalidParams
	}
	s.bucket = args[1]

	for i := 2; i < num; i += 2 {
		s.val[args[i]] = args[i+1]
	}
	return nil
}
