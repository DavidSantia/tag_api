package tag_api

import (
	"github.com/jmoiron/sqlx"
)

// DB loaders

func (bs *BoltService) loadGroups() {
	var err error
	var query string
	var g Group
	var rows *sqlx.Rows

	// Query groups
	query = makeQuery(g, GroupQuery)
	Log.Debug.Printf("GroupQuery: %s\n", query)
	rows, err = bs.ds.db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load Groups: %v\n", err)
		return
	}

	// Load into GroupMap
	bs.GroupMap = make(GroupMap)
	for rows.Next() {
		g = Group{
			ImagesGroupsMap: make(ImagesGroupsMap),
		}
		err = rows.StructScan(&g)
		if err != nil {
			Log.Error.Printf("Load Group: %v\n", err)
			continue
		}
		bs.GroupMap[g.Id] = g
	}
	Log.Info.Printf("Load Groups: %d entries total\n", len(bs.GroupMap))
}

func (bs *BoltService) loadImagesGroups() {
	var err error
	var query string
	var g Group
	var ig ImagesGroups
	var rows *sqlx.Rows
	var ok bool
	var entries, ignored int

	// Query group-image mapping
	query = makeQuery(ig, ImagesGroupsQuery)
	Log.Debug.Printf("ImagesGroupsQuery: %s\n", query)
	rows, err = bs.ds.db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load ImagesGroups: %v\n", err)
		return
	}

	// Load into GroupMap
	for rows.Next() {
		err = rows.StructScan(&ig)
		if err != nil {
			Log.Error.Printf("Load ImagesGroups: %v\n", err)
			continue
		}
		_, ok = bs.ImageMap[ig.ImageId]
		if !ok {
			// Skip any inage that is not in ImageMap
			ignored++
			continue
		}
		g, ok = bs.GroupMap[ig.GroupId]
		if !ok {
			Log.Error.Printf("Load ImagesGroups: ImageId %d on invalid GroupId %d\n", ig.ImageId, ig.GroupId)
		}
		g.ImagesGroupsMap[ig.ImageId] = true
		bs.GroupMap[ig.GroupId] = g
		entries++
	}
	if ignored > 0 {
		Log.Info.Printf("Load ImagesGroups: %d entries total [ignored %d invalid ImageIds]\n",
			entries, ignored)
	} else {
		Log.Info.Printf("Load ImagesGroups: %d entries total\n", entries)
	}
}
