package cmd

import (
	"boltview/pkg/boltdb"
	"fmt"
)

const (
	cmdBuckets         = "buckets"
	descriptionBuckets = "show buckets"
)

func init() {
	RegisterCommand(cmdBuckets, &buckets{})
}

type buckets struct {
	base
	filter  []string
	buckets []string
}

func (b *buckets) Open(opts ...ParseOption) (Command, error) {
	o := &buckets{base: base{
		name:        cmdBuckets,
		cmd:         cmdBuckets,
		description: descriptionBuckets,
	}}
	for _, opt := range opts {
		opt(o)
	}
	return o, nil
}

func (b *buckets) exec() error {
	buckets, err := boltdb.Buckets()
	if err != nil {
		return err
	}
	b.buckets = buckets
	return nil
}

func (b *buckets) parse(args []string) error {
	if len(b.params) <= 1 {
		return nil
	}
	b.params = args
	b.filter = b.params[1:]
	return nil
}

func (b *buckets) ok() {
	fmt.Println(b.buckets)
}
