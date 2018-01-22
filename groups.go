package tag_api

import (
	"github.com/jmoiron/sqlx"
)

// DB loaders

func (data *ApiData) LoadGroups() {
	var err error
	var query string
	var g Group
	var rows *sqlx.Rows

	data.GroupMap = make(GroupMap)

	// Load partner map
	query = data.MakeQuery(g, GroupQuery)
	Log.Debug.Printf("GroupQuery: %s\n", query)
	rows, err = data.Db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load Groups: %v\n", err)
		return
	}
	for rows.Next() {
		g = Group{
			ImagesGroupsMap: make(ImagesGroupsMap),
		}
		err = rows.StructScan(&g)
		if err != nil {
			Log.Error.Printf("Load Group: %v\n", err)
			continue
		}
		data.GroupMap[g.Id] = g
	}
	Log.Info.Printf("Load Partners: %d entries total\n", len(data.GroupMap))
}

func (data *ApiData) LoadImagesGroups() {
	var err error
	var query string
	var g Group
	var ig ImagesGroups
	var rows *sqlx.Rows
	var ok bool
	var entries, ignored int

	// Get partner merchant mapping
	query = data.MakeQuery(ig, ImagesGroupsQuery)
	Log.Debug.Printf("ImagesGroupsQuery: %s\n", query)
	rows, err = data.Db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load ImagesGroups: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.StructScan(&ig)
		if err != nil {
			Log.Error.Printf("Load ImagesGroups: %v\n", err)
			continue
		}
		_, ok = data.ImageMap[ig.ImageId]
		if !ok {
			// Skip any inage that is not in ImageMap
			ignored++
			continue
		}
		g, ok = data.GroupMap[ig.GroupId]
		if !ok {
			Log.Error.Printf("Load ImagesGroups: ImageId %d on invalid GroupId %d\n", ig.ImageId, ig.GroupId)
		}
		g.ImagesGroupsMap[ig.ImageId] = true
		data.GroupMap[ig.GroupId] = g
		entries++
	}
	if ignored > 0 {
		Log.Info.Printf("Load ImagesGroups: %d entries total [ignored %d invalid ImageIds]\n",
			entries, ignored)
	} else {
		Log.Info.Printf("Load ImagesGroups: %d entries total\n", entries)
	}
}
