package tag_api

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/jmoiron/sqlx"
)

// DB loaders

func (data *ApiData) UpdateImages() {
	var err error
	var query string
	var image Image
	var rows *sqlx.Rows

	imageMap := make(ImageMap)

	// Load images
	query = data.MakeQuery(image, ImageQuery)
	Log.Debug.Printf("ImageQuery: %s\n", query)
	rows, err = data.Db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Update Images: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.StructScan(&image)
		if err != nil {
			Log.Error.Printf("Update Images: %v\n", err)
			continue
		}
		imageMap[image.Id] = image
	}
	Log.Info.Printf("Update Images: %d entries fetched from MySQL\n", len(imageMap))

	// Store in Bolt
	err = data.BoltDb.Update(func(tx *bolt.Tx) (e error) {
		var bucket *bolt.Bucket
		var b bytes.Buffer
		var k []byte

		bucket, e = tx.CreateBucketIfNotExists(data.BoltBucket)
		if e != nil {
			e = fmt.Errorf("Create bucket %s in Bolt: %v", data.BoltBucket, e)
			return
		}
		k = []byte("Images")

		enc := gob.NewEncoder(&b)
		e = enc.Encode(imageMap)
		if e != nil {
			e = fmt.Errorf("Encode ImageMap for Bolt: %v", e)
			return
		}

		e = bucket.Put(k, b.Bytes())
		if e != nil {
			e = fmt.Errorf("Store ImageMap in Bolt: %v", e)
			return
		}
		return
	})
	Log.Info.Printf("Update Images: %d entries stored in Bolt\n", len(imageMap))
}
