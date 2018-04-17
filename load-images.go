package tag_api

import (
	"github.com/jmoiron/sqlx"
)

func (data *ApiData) LoadImages() {
	var err error
	var query string
	var image Image
	var rows *sqlx.Rows

	data.ImageMap = make(ImageMap)

	// Load images
	query = data.MakeQuery(image, ImageQuery)
	Log.Debug.Printf("ImageQuery: %s\n", query)
	rows, err = data.Db.Queryx(query)
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
		data.ImageMap[image.Id] = image
	}
	Log.Info.Printf("Load Images: %d entries total\n", len(data.ImageMap))
}
