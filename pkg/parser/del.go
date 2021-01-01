package parser

import (
	"boltview/pkg/boltdb"
	"strings"
)

const (
	cmdDel         = "del"
	descriptionDel = "delete buckets"
)

type del struct {
	base
}

func init() {
	register(newDel())
}

func newDel() *del {
	return &del{
		base: base{
			name:        cmdDel,
			cmd:         cmdDel,
			description: descriptionDel,
		}}
}

func (d *del) parse(args []string) error {
	if len(args) < 2 {
		return ErrInvalidParams
	}
	d.params = args
	return nil
}

func (d *del) exec() error {
	for _, p := range d.params[1:] {
		if strings.Contains(p, ".") {
			args := strings.Split(p, ".")
			err := boltdb.DeleteKey(args[0], args[1])
			if err != nil {
				return nil
			}
		} else {
			err := boltdb.DeleteBucket(p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
