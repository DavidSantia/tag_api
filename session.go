package tag_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/julienschmidt/httprouter"
)

func (data *ApiData) InitSessions() {

	// Initialise the session manager
	data.sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

	// maximum length of time a session can be inactive
	data.sessionManager.IdleTimeout(30 * time.Minute)

	// maximum length of time that a session is valid
	data.sessionManager.Lifetime(8 * time.Hour)

	// whether the session cookie should be retained after a user closes their browser
	data.sessionManager.Persist(false)
}

// HTTP Handlers

func makeHandleAuthenticate(ds *DbService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var user User
		var b []byte

		auth := r.Header.Get("Authorization")
		if len(auth) == 0 || !strings.Contains(auth, "Bearer") {
			HandleError(w, http.StatusBadRequest, r.RequestURI, fmt.Errorf("Invalid authentication: %s", auth))
			return
		}

		// Decode payload from JSON Web Token
		token := strings.TrimSpace(strings.Replace(auth, "Bearer", "", 1))
		payload, headers, err := jose.Decode(token, JwtKey)
		if err != nil {
			HandleError(w, http.StatusBadRequest, r.RequestURI, err)
			return
		}

		Log.Debug.Printf("Headers = %+v\n", headers)
		Log.Debug.Printf("Payload = %s\n", payload)

		// Make sure payload is JSON
		if len(payload) == 0 || payload[0] != '{' {
			HandleError(w, http.StatusBadRequest, r.RequestURI, fmt.Errorf("Payload decryption failed"))
			return
		}

		Log.Debug.Printf("Authenticate Payload: %s\n", payload)

		pl := JwtPayload{}
		err = json.Unmarshal([]byte(payload), &pl)
		if err != nil {
			HandleError(w, http.StatusBadRequest, r.RequestURI, err)
			return
		}

		// Lookup user by id
		user, err = userFind(ds, pl)
		if err != nil {
			HandleError(w, http.StatusBadRequest, r.RequestURI, err)
			return
		}

		// Store user data in session
		session := d.sessionManager.Load(r)
		err = session.PutInt64(w, "gid", user.GroupId)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}
		err = session.PutInt64(w, "uid", user.Id)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}

		// Reply
		b, err = json.Marshal(user)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		HandleReply(w, http.StatusOK, string(b)+"\n")
	}
}

func makeHandleAuthKeepAlive(cs ContentService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		session := d.sessionManager.Load(r)
		id, err := session.GetInt64("uid")
		if err != nil || id == 0 {
			HandleError(w, http.StatusUnauthorized, r.RequestURI, fmt.Errorf("Session not authenticated"))
			return
		}

		// Reset inactivity period
		err = session.Touch(w)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}

		// Reply
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		HandleReply(w, http.StatusOK, `{"message":"IdleTimeout Updated","status":"OK"}`)
	}
}

// Data Interfaces

func GetGroupIdFromSession(r *http.Request) (gid int64, err error) {

	// See if session is authenticated
	session := d.sessionManager.Load(r)
	gid, err = session.GetInt64("gid")
	if err != nil || gid == 0 {
		err = fmt.Errorf("Session not authenticated")
	}
	return
}

func GetUserFromSession(ds *DbService, r *http.Request) (user User, err error) {
	var ok bool

	// See if session is authenticated
	session := d.sessionManager.Load(r)
	id, err := session.GetInt64("uid")
	if err != nil || id == 0 {
		err = fmt.Errorf("Session not authenticated")
		return
	}

	// Lookup user
	user, ok = ds.GetUser(id)
	if !ok {
		err = fmt.Errorf("UserId %d not valid", id)
	}
	return
}

func userFind(ds *DbService, pl JwtPayload) (user User, err error) {
	var ok bool

	// Validate payload
	if pl.UserId == 0 {
		err = fmt.Errorf("Invalid UserId in payload")
		return
	}
	if len(pl.Guid) == 0 {
		err = fmt.Errorf("Invalid Guid in payload")
		return
	}

	// Lookup user
	user, ok = ds.GetUser(pl.UserId)
	if !ok {
		err = fmt.Errorf("UserId %d not valid", pl.UserId)
		return
	}

	// Validate Guid
	if user.Guid != pl.Guid {
		err = fmt.Errorf("Invalid Guid %s specified for UserId %d", pl.Guid, pl.UserId)
		return
	}

	// Validate Status
	if !user.Status {
		err = fmt.Errorf("Expired User %d", pl.UserId)
	}
	return
}
