package tag_api

import (
	"github.com/jmoiron/sqlx"
)

func (bs *BoltService) loadImages() {
	var err error
	var query string
	var image Image
	var rows *sqlx.Rows

	bs.ImageMap = make(ImageMap)

	// Load images
	query = makeQuery(image, ImageQuery)
	Log.Debug.Printf("ImageQuery: %s\n", query)
	rows, err = bs.db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load Images: %v\n", err)
		return
	}
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
