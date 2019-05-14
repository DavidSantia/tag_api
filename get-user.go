package tag_api

import (
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent"
)

func (ds *DbService) GetUser(id int64, txn newrelic.Transaction) (user User, ok bool) {
	var err error
	var rows *sqlx.Rows

	segment := UserSegment
	if txn != nil {
		segment.StartTime = newrelic.StartSegmentNow(txn)
		defer segment.End()
	}

	// Query users
	Log.Debug.Printf("UserQuery: %s\n", segment.ParameterizedQuery)
	segment.QueryParameters = map[string]interface{}{"id": id}
	rows, err = ds.Queryx(segment.ParameterizedQuery, id)
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
