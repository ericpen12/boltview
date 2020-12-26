package exec

import "boltview/boltdb"

const (
	cmdDel         = "del"
	descriptionDel = "delete buckets"
)

type del struct {
	create
}

func init() {
	register(newDel())
}

func newDel() *del {
	return &del{
		create{
			base: base{
				name:        cmdDel,
				cmd:         cmdDel,
				description: descriptionDel,
			}}}
}

func (d *del) Exec() error {
	for _, bucket := range d.bucketNames {
		err := boltdb.DeleteBucket(bucket)
		if err != nil {
			return err
		}
	}
	return nil
}
