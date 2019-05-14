package tag_api

import (
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent"
)

func (bs *BoltService) loadImages(ds *DbService, txn newrelic.Transaction) {
	var err error
	var image Image
	var rows *sqlx.Rows

	if txn != nil {
		ImageSegment.StartTime = newrelic.StartSegmentNow(txn)
		defer ImageSegment.End()
	}

	// Query images
	Log.Debug.Printf("ImageQuery: %s\n", ImageSegment.ParameterizedQuery)
	rows, err = ds.Queryx(ImageSegment.ParameterizedQuery)
	if err != nil {
		Log.Error.Printf("Load Images: %v\n", err)
		return
	}

	// Load into ImageMap
	bs.ImageMap = make(ImageMap)
	for rows.Next() {
		err = rows.StructScan(&image)
		if err != nil {
			Log.Error.Printf("Load Images: %v\n", err)
			continue
		}
		bs.ImageMap[image.Id] = image
	}
	Log.Info.Printf("Load Images: %d entries total\n", len(bs.ImageMap))
}
