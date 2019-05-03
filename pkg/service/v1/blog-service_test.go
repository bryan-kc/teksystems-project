package v1

import (
	"github.com/boltdb/bolt"
	"log"
	"testing"
)

func Test_BlogServer_Create(t *testing.T) {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Fatalf("An error occurred while setting up the database: %s", err)
	}

	defer db.Close()

	s := NewBlogServiceServer(db)

}