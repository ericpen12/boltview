package exec

import (
	"boltview/boltdb"
	"errors"
)

const (
	cmdSet         = "set"
	descriptionSet = "create a specific k/v"
)

var (
	ErrInvalidParams = errors.New("invalid params")
)

type set struct {
	base
	bucket string
	val    map[string]string
}

func init() {
	register(newSet())
}

func newSet() *set {
	return &set{base: base{
		name:        cmdSet,
		cmd:         cmdSet,
		description: descriptionSet,
	}}
}

func (s *set) Error(err error) {
	print(err)
}

func (s *set) Exec() error {
	for k, v := range s.val {
		err := boltdb.Set(s.bucket, k, []byte(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *set) Parse(args []string) error {
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
