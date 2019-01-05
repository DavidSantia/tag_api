package tag_api

import (
	"github.com/jmoiron/sqlx"
)

func (bs *BoltService) loadUsers() {
	var err error
	var query string
	var user User
	var rows *sqlx.Rows

	// Query users
	query = makeQuery(user, UserQuery)
	Log.Debug.Printf("UserQuery: %s\n", query)
	rows, err = bs.ds.db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load Users: %v\n", err)
		return
	}

	// Load into UserMap
	bs.UserMap = make(UserMap)
	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			Log.Error.Printf("Load Users: %v\n", err)
			continue
		}
		bs.UserMap[user.Id] = user
	}
	Log.Info.Printf("Load Users: %d entries total\n", len(bs.UserMap))
}
