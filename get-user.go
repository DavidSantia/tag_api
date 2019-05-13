package tag_api

import (
	"github.com/jmoiron/sqlx"
)

func (ds *DbService) GetUser(id int64) (user User, ok bool) {
	var err error
	var query string
	var rows *sqlx.Rows

	// Query users
	query = makeQuery(user, UserQuery, id)
	Log.Debug.Printf("UserQuery: %s\n", query)
	rows, err = ds.Queryx(query)
	if err != nil {
		Log.Error.Printf("Get User: %v\n", err)
		return
	}

	// Load into UserMap
	for rows.Next() {
		err = rows.StructScan(&user)
		ok = err == nil
		return
	}
	return
}
