package parser

import "boltview/pkg/boltdb"

const (
	cmdBuckets         = "buckets"
	descriptionBuckets = "show buckets"
)

type buckets struct {
	base
	filter  []string
	buckets []string
}

func init() {
	register(newBuckets())
}

func newBuckets() *buckets {
	return &buckets{base: base{
		name:        cmdBuckets,
		cmd:         cmdBuckets,
		description: descriptionBuckets,
	}}
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
	writeToConsole(b.buckets)
}
