package tag_api

import (
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent"
)

// DB loaders

func (bs *BoltService) loadGroups(ds *DbService, txn newrelic.Transaction) {
	var err error
	var g Group
	var rows *sqlx.Rows

	if txn != nil {
		GroupSegment.StartTime = newrelic.StartSegmentNow(txn)
		defer GroupSegment.End()
	}

	// Query groups
	Log.Debug.Printf("GroupQuery: %s\n", GroupSegment.ParameterizedQuery)
	rows, err = ds.Queryx(GroupSegment.ParameterizedQuery)
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

func (bs *BoltService) loadImagesGroups(ds *DbService, txn newrelic.Transaction) {
	var err error
	var g Group
	var ig ImagesGroups
	var rows *sqlx.Rows
	var ok bool
	var entries, ignored int

	if txn != nil {
		ImageGroupSegment.StartTime = newrelic.StartSegmentNow(txn)
		defer ImageGroupSegment.End()
	}

	// Query group-image mapping
	Log.Debug.Printf("ImagesGroupsQuery: %s\n", ImageGroupSegment.ParameterizedQuery)
	rows, err = ds.Queryx(ImageGroupSegment.ParameterizedQuery)
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
