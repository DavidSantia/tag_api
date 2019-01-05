package tag_api

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/boltdb/bolt"
)

// Store current content to BoltDb

func (bs *BoltService) storeUsers() {
	var err error

	// Store in Bolt
	err = bs.boltDb.Update(func(tx *bolt.Tx) (e error) {
		var bucket *bolt.Bucket
		var b bytes.Buffer
		var k []byte

		bucket, e = tx.CreateBucketIfNotExists(bs.settings.boltBucket)
		if e != nil {
			e = fmt.Errorf("Create bucket %s in Bolt: %v", bs.settings.boltBucket, e)
			return
		}
		k = []byte("users")

		enc := gob.NewEncoder(&b)
		e = enc.Encode(bs.UserMap)
		if e != nil {
			e = fmt.Errorf("Encode UserMap for Bolt: %v", e)
			return
		}

		e = bucket.Put(k, b.Bytes())
		if e != nil {
			e = fmt.Errorf("Store UserMap in Bolt: %v", e)
			return
		}
		return
	})
	if err != nil {
		Log.Error.Printf("Store Users: %v", err)
		return
	}

	Log.Info.Println("Store Users: Wrote UserMap to Bolt")
}
