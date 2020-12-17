package boltdb

import (
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
	tests := []struct {
		bucket string
		result error
	}{
		{"test", nil},
		{"test", ErrBucketExist},
		{"123", nil},
	}
	for _, test := range tests {
		err := CreateBucket(test.bucket)
		if err != test.result {
			t.Fatalf("bucket name is: %s, the excepted is: %v, but the actual is: %v",
				test.bucket, test.result, err)
		}
	}

}

func TestDeleteBucket(t *testing.T) {
	err := DeleteBucket("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteKey(t *testing.T) {
	err := DeleteKey("test", "name")
	if err != nil {
		t.Error(err)
	}
}
