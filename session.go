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
	data.SessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

	// maximum length of time a session can be inactive
	data.SessionManager.IdleTimeout(30 * time.Minute)

	// maximum length of time that a session is valid
	data.SessionManager.Lifetime(8 * time.Hour)

	// whether the session cookie should be retained after a user closes their browser
	data.SessionManager.Persist(false)
}

// HTTP Handlers

func HandleAuthenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u User
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
	u, err = d.UserFind(pl)
	if err != nil {
		HandleError(w, http.StatusBadRequest, r.RequestURI, err)
		return
	}

	// Send message to content server
	uMsg := UserMessage{
		Command:   "adduser",
		Id:        u.Id,
		GroupId:   u.GroupId,
		Guid:      u.Guid,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
	b, err = json.Marshal(uMsg)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
		return
	}
	d.Nconn.Publish("users", b)
	Log.Info.Printf("Authenticate: %s %s [id=%d]\n", u.FirstName, u.LastName, u.Id)

	// Store user data in session
	session := d.SessionManager.Load(r)
	err = session.PutInt64(w, "gid", u.GroupId)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
		return
	}
	err = session.PutInt64(w, "uid", u.Id)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
		return
	}

	// Reply
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	HandleReply(w, http.StatusOK, string(b)+"\n")
}

func HandleAuthKeepAlive(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := d.SessionManager.Load(r)
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

// Data Interfaces

func GetGroupIdFromSession(r *http.Request) (gid int64, err error) {

	// See if session is authenticated
	session := d.SessionManager.Load(r)
	gid, err = session.GetInt64("gid")
	if err != nil || gid == 0 {
		err = fmt.Errorf("Session not authenticated")
	}
	return
}

func GetUserFromSession(r *http.Request) (u User, err error) {
	var ok bool

	// See if session is authenticated
	session := d.SessionManager.Load(r)
	id, err := session.GetInt64("uid")
	if err != nil || id == 0 {
		err = fmt.Errorf("Session not authenticated")
		return
	}

	// Lookup user
	u, ok = d.UserMap[id]
	if !ok {
		err = fmt.Errorf("UserId %d not valid", id)
	}
	return
}

func (data *ApiData) UserFind(pl JwtPayload) (u User, err error) {
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
	u, ok = data.UserMap[pl.UserId]
	if !ok {
		err = fmt.Errorf("UserId %d not valid", pl.UserId)
		return
	}

	// Validate Guid
	if u.Guid != pl.Guid {
		err = fmt.Errorf("Invalid Guid %s specified for UserId %d", pl.Guid, pl.UserId)
		return
	}

	// Validate Status
	if !u.Status {
		err = fmt.Errorf("Expired User %d", pl.UserId)
	}
	return
}
