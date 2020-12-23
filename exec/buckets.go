package exec

import "boltview/boltdb"

const (
	cmdBuckets         = "buckets"
	descriptionBuckets = "show buckets"
)

type buckets struct {
	base
	filter  []string
	buckets []string
}

func (b *buckets) CommandName() string {
	return b.name
}

func (b *buckets) Description() string {
	return b.description
}

func (b *buckets) Error(err error) {
	print(err)
}

func (b *buckets) Exec() error {
	buckets, err := boltdb.Buckets()
	if err != nil {
		return err
	}
	b.buckets = buckets
	return nil
}

func (b *buckets) Parse(args []string) error {
	if len(b.options) <= 1 {
		return nil
	}
	b.options = args
	b.filter = b.options[1:]
	return nil
}

func (b *buckets) Ok() {
	print(b.buckets)
}

func NewCmdBuckets() *buckets {
	return &buckets{base{
		name:        cmdBuckets,
		cmd:         cmdBuckets,
		description: descriptionBuckets,
		options:     nil,
	}, nil, nil}
}
