package tag_api

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/boltdb/bolt"
)

// DB loaders

func (data *ApiData) LoadImages() {
	var err error

	// Load images
	err = data.BoltDb.View(func(tx *bolt.Tx) (e error) {
		var bucket *bolt.Bucket
		var v, k []byte

		bucket = tx.Bucket(data.BoltBucket)
		if bucket == nil {
			e = fmt.Errorf("Bolt bucket %s not found", data.BoltBucket)
			return
		}
		k = []byte("Images")

		v = bucket.Get(k)
		if v == nil {
			e = fmt.Errorf("Bolt key %s not found", k)
			return
		}

		b := bytes.NewBuffer(v)
		dec := gob.NewDecoder(b)

		e = dec.Decode(&data.ImageMap)
		if e != nil {
			e = fmt.Errorf("Parse ImageMap gob from Bolt: %v", e)
			return
		}
		return
	})

	if err != nil {
		Log.Error.Printf("Load Images: %v\n", err)
	}

	Log.Info.Printf("Load Images: %d entries loaded from Bolt\n", len(data.ImageMap))
}
