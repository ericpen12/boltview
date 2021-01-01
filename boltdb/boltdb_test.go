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
	bucket := "bucket_test_keys"
	err := CreateBucket(bucket)
	if err != nil {
		t.Error(err)
	}
	Convey("test: set (normal test)", t, func() {
		keys, err := Keys(bucket)
		So(keys, ShouldBeNil)
		So(err, ShouldBeNil)

		err = Set(bucket, "k", []byte("v"))
		if err != nil {
			t.Error(err)
		}

		keys, err = Keys(bucket)
		So(keys, ShouldNotBeNil)
		So(err, ShouldBeNil)
		err = DeleteBucket(bucket)
		if err != nil {
			t.Error(err)
		}
	})

	Convey("test: set (if bucket does not exist", t, func() {
		keys, err := Keys("bucket_not_exist_test_keys")
		So(keys, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
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
	b := "bucket111"
	err := CreateBucket(b)
	if err != nil {
		t.Fatalf("create bucket error, %v", err)
	}
	Convey("set key (normal test)", t, func() {
		So(Set(b, "k1", []byte("v1")), ShouldBeNil)
		So(Set(b, "k1", []byte("v2")), ShouldBeNil)
		So(Set(b, "k1", []byte("")), ShouldBeNil)
	})

	Convey("set key (special test)", t, func() {
		So(Set(b, "", []byte("v1")), ShouldNotBeNil)
		So(Set(b, "k1", nil), ShouldNotBeNil)
	})
	err = DeleteBucket(b)
	if err != nil {
		t.Fatalf("delete bucket error: %v", err)
	}

	Convey("set key (bucket does not exist)", t, func() {
		So(Set("notSuchBucket", "k1", []byte("v1")), ShouldEqual, ErrBucketNotExist)
	})
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
	Convey("test delete bucket", t, func() {
		_ = CreateBucket("test")
		So(DeleteBucket("test"), ShouldBeNil)
		So(DeleteBucket("bucket1222121"), ShouldBeError)
		So(DeleteBucket(""), ShouldEqual, bbolt.ErrBucketNameRequired)
	})
}

func TestBuckets(t *testing.T) {
	Convey("test lookup bucket", t, func() {
		bucketList := []string{"lookB", "lookB2"}

		// create bucket
		for _, v := range bucketList {
			_ = CreateBucket(v)
		}
		result, err := Buckets()
		So(err, ShouldBeNil)

		// check the specific bucket if it does exist
		for _, b := range bucketList {
			So(result, ShouldContain, b)
		}

		// delete created buckets
		for _, v := range bucketList {
			_ = DeleteBucket(v)
		}
	})
}

func TestDeleteKey(t *testing.T) {
	err := DeleteKey("test", "name")
	if err != nil {
		t.Error(err)
	}
}
