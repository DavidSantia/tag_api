package tag_api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HTTP Handlers

func makeHandleUser(cs ContentService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var err error
		var user User

		user, err = GetUserFromSession(cs, r)
		if err != nil {
			HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
			return
		}

		var b []byte
		b, err = json.Marshal(user)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}

		HandleReply(w, http.StatusOK, string(b))
	}
}

func (bs *BoltService) addUser(msg []byte) (err error) {
	var user User

	// Add user
	err = json.Unmarshal(msg, &user)
	if err != nil {
		return
	}
	bs.UserMap[user.Id] = user

	Log.Info.Printf("Add User: %s %s [id=%d]\n", user.FirstName, user.LastName, user.Id)
	return
}
