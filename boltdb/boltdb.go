package boltdb

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

var db *bolt.DB

func Open(path string) {
	var err error
	db, err = bolt.Open(path, 0666, nil)
	if err != nil {
		log.Println("Cannot connect db, path: ", path)
	}
	log.Println("connected!")
}

func Keys(bucket string) ([]string, error) {
	var keys []string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
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
