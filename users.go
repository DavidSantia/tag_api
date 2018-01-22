package tag_api

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

// DB loaders

func (data *ApiData) LoadUsers() {
	var err error
	var query string
	var u User
	var rows *sqlx.Rows

	data.UserMap = make(UserMap)

	// Load images
	query = data.MakeQuery(u, UserQuery)
	Log.Debug.Printf("UserQuery: %s\n", query)
	rows, err = data.Db.Queryx(query)
	if err != nil {
		Log.Error.Printf("Load Users: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.StructScan(&u)
		if err != nil {
			Log.Error.Printf("Load Users: %v\n", err)
			continue
		}
		data.UserMap[u.Id] = u
	}
	Log.Info.Printf("Load Users: %d entries total\n", len(data.UserMap))
}

// HTTP Handlers

func HandleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u, err := GetUserFromSession(r)
	if err != nil {
		HandleError(w, http.StatusUnauthorized, r.RequestURI, err.Error())
		return
	}

	var b []byte
	b, err = json.Marshal(u)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err.Error())
		return
	}

	HandleReply(w, http.StatusOK, string(b))
}
