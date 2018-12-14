package postbox

import (
	"github.com/boltdb/bolt"
	"path"
	"time"
)

type BoltClient struct {
	Path string
	db *bolt.DB
}

func NewBoltClient(path string) *BoltClient {
	c := &BoltClient{Path: path}
	return c
}

func (c *BoltClient) Open() error {
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	c.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("postbox")); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (c *BoltClient) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func defaultBoltClient() *BoltClient {
	return NewBoltClient(path.Join(getPostboxDir(), "postbox.db"))
}

var boltClient = defaultBoltClient()

