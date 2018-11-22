package postbox

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
)

//go:generate protoc --go_out=. postbox.proto

func Get(id uint64) (*Letter, error) {
	boltClient.Open()
	defer boltClient.Close()

	var letter *Letter
	err := boltClient.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("postbox"))
		dbVal := bucket.Get(itob(id))
		if dbVal == nil {
			return fmt.Errorf("cannot get from db with id: %d", id)
		}
		var err error
		letter, err = unmarshallLetter(dbVal)
		if err != nil {
			return fmt.Errorf("failed to unmarshall letter: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read db: %v", err)
	}

	return letter, nil
}

func List() ([]*Letter, error) {
	boltClient.Open()
	defer boltClient.Close()

	var letters []*Letter
	err := boltClient.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("postbox"))
		err := bucket.ForEach(func(k, v []byte) error {
			letter, err := unmarshallLetter(v)
			if err != nil {
				return fmt.Errorf("failed to unmarshall letter: %v", err)
			}
			letters = append(letters, letter)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read db: %v", err)
	}

	return letters, nil
}

func Clear() error {
	boltClient.Open()
	defer boltClient.Close()

	err := boltClient.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("postbox"))
		err := bucket.ForEach(func(k, v []byte) error {
			err := bucket.Delete(k)
			if err != nil {
				return fmt.Errorf("failed to delete letter: %v", err)
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update db: %v", err)
	}
	return nil
}

func marshalLetter(letter *Letter) ([]byte, error) {
	return proto.Marshal(letter)
}

func unmarshallLetter(data []byte) (*Letter, error) {
	var res Letter
	err := proto.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
