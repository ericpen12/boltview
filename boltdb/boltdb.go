package boltdb

import (
	"errors"
	"fmt"
	"github.com/coreos/bbolt"
	bolt "go.etcd.io/bbolt"
	"log"
	"os"
	"strings"
	"time"
)

var db *bolt.DB

const (
	defaultTimeout = time.Second * 3
)

var (
	ErrBucketExist    = errors.New("bucket already exists")
	ErrBucketNotExist = errors.New("bucket does not exist")
)

func Open(path string) {
	var err error
	db, err = bolt.Open(path, 0666, &bolt.Options{Timeout: defaultTimeout})
	if err != nil {
		log.Printf("Cannot connect db: %s, err: %v", path, err)
		os.Exit(0)
	}
	log.Println("connected!")
}

func Keys(bucket string) ([]string, error) {
	var keys []string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New(fmt.Sprintf("bucket does not exist: %s", bucket))
		}
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			keys = append(keys, string(k))
		}
		return nil

	})
	return keys, err
}

func Buckets() ([]string, error) {
	var buckets []string
	err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			buckets = append(buckets, string(name))
			return nil
		})
	})
	return buckets, err
}

func Get(field string) ([]byte, error) {
	var bucket, key string
	var value []byte
	s := strings.Split(field, ".")
	bucket = s[0]
	if len(s) > 1 {
		key = s[1]
	}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value = append(value, b.Get([]byte(key))...)
		return nil
	})
	return value, err
}

func Set(bu, key string, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bu))
		if err := b.Put([]byte(key), value); err != nil {
			return err
		}
		return nil
	})
}

func CreateBucket(bucket string) error {
	if len(bucket) == 0 {
		return bbolt.ErrBucketNameRequired
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		if err == bbolt.ErrBucketExists {
			return ErrBucketExist
		}
		if err != nil {
			return err
		}
		return nil
	})
}

func DeleteBucket(bucket string) error {
	if len(bucket) == 0 {
		return bbolt.ErrBucketNameRequired
	}
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucket))
	})
}

func DeleteKey(bucket, key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return ErrBucketNotExist
		}
		return b.Delete([]byte(key))
	})
}
