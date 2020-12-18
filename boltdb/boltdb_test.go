package boltdb

import (
	"github.com/coreos/bbolt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func init() {
	Open("../db/test.db")
}
func TestKeys(t *testing.T) {
	keys, err := Keys("test")
	if err != nil {
		t.Fatal("get keys err, ", err)
	}
	t.Log(keys)
}

func TestBuckets(t *testing.T) {
	b, err := Buckets()
	if err != nil {
		t.Fatal("get buckets error, ", err)
	}
	t.Log(b)
}

func TestGet(t *testing.T) {
	b, err := Get("test.name")
	if err != nil {
		t.Fatal(err)
	}
	if b == nil {
		t.Fatal("get empty")
	}
	ioutil.WriteFile("pay.json", b, 0700)
	t.Log(string(b))
}

func TestSet(t *testing.T) {
	err := Set("test", "name", []byte("tom"))
	if err != nil {
		t.Error(err)
	}
}

func TestCreateBucket(t *testing.T) {
	Convey("test create bucket", t, func() {
		So(CreateBucket("bucket1"), ShouldBeNil)
		So(CreateBucket("1bucket1"), ShouldBeNil)
		So(CreateBucket("中文"), ShouldBeNil)
		_ = DeleteBucket("bucket1")
		_ = DeleteBucket("1bucket1")
		_ = DeleteBucket("中文")
		So(CreateBucket(""), ShouldEqual, bbolt.ErrBucketNameRequired)
	})

}

func TestDeleteBucket(t *testing.T) {
	Convey("test create bucket", t, func() {
		_ = CreateBucket("test")
		So(DeleteBucket("test"), ShouldBeNil)
		So(DeleteBucket("bucket1222121"), ShouldBeError)
		So(DeleteBucket(""), ShouldEqual, bbolt.ErrBucketNameRequired)
	})
}

func TestDeleteKey(t *testing.T) {
	err := DeleteKey("test", "name")
	if err != nil {
		t.Error(err)
	}
}
