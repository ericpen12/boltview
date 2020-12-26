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
