package tag_api

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/boltdb/bolt"
)

// Refresh data loaded from Bolt Db

func (bs *BoltService) refreshUsers() {
	var err error

	// Refresh users
	err = bs.boltDb.View(func(tx *bolt.Tx) (e error) {
		var bucket *bolt.Bucket
		var v, k []byte

		bucket = tx.Bucket(bs.settings.boltBucket)
		if bucket == nil {
			e = fmt.Errorf("Bolt bucket %s not found", bs.settings.boltBucket)
			return
		}
		k = []byte("users")

		v = bucket.Get(k)
		if v == nil {
			e = fmt.Errorf("Bolt key %s not found", k)
			return
		}

		b := bytes.NewBuffer(v)
		dec := gob.NewDecoder(b)

		e = dec.Decode(&bs.UserMap)
		if e != nil {
			e = fmt.Errorf("Parse UserMap gob from Bolt: %v", e)
			return
		}
		return
	})

	if err != nil {
		Log.Error.Printf("Refresh Users: %v\n", err)
	}

	Log.Info.Printf("Refresh Users: %d entries loaded from Bolt\n", len(bs.UserMap))
}
