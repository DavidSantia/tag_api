package tag_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
		HandleError(w, http.StatusBadRequest, r.RequestURI, "Invalid authentication: "+auth)
		return
	}

	// Decode payload from JSON Web Token
	token := strings.TrimSpace(strings.Replace(auth, "Bearer", "", 1))
	payload, headers, err := jose.Decode(token, JwtKey)
	if err != nil {
		HandleError(w, http.StatusBadRequest, r.RequestURI, err.Error())
		return
	}

	Log.Debug.Printf("Headers = %+v\n", headers)
	Log.Debug.Printf("Payload = %s\n", payload)

	// Make sure payload is JSON
	if len(payload) == 0 || payload[0] != '{' {
		HandleError(w, http.StatusBadRequest, r.RequestURI, "Payload decryption failed")
		return
	}

	Log.Debug.Printf("Authenticate Payload: %s\n", payload)

	pl := JwtPayload{}
	err = json.Unmarshal([]byte(payload), &pl)
	if err != nil {
		HandleError(w, http.StatusBadRequest, r.RequestURI, err.Error())
		return
	}

	// Lookup user by id
	u, err = d.UserFind(pl)
	if err != nil {
		HandleError(w, http.StatusBadRequest, r.RequestURI, err.Error())
		return
	}
	b, err = json.Marshal(u)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err.Error())
		return
	}

	// Store user data in session
	gid := fmt.Sprintf("%d", u.GroupId)

	session := d.SessionManager.Load(r)
	err = session.PutString(w, "gid", gid)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err.Error())
		return
	}
	err = session.PutBytes(w, "json", b)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err.Error())
		return
	}

	// Reply
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	HandleReply(w, http.StatusOK, string(b)+"\n")
}

// Data Interfaces

func GetGroupIdFromSession(r *http.Request) (group_id int64, err error) {
	var gid string

	// See if session is authenticated
	session := d.SessionManager.Load(r)
	gid, err = session.GetString("gid")
	if len(gid) == 0 {
		err = fmt.Errorf("Session not authenticated")
		return
	}
	group_id, err = strconv.ParseInt(gid, 10, 64)
	return
}

func GetUserFromSession(r *http.Request) (u User, err error) {

	// See if session is authenticated
	session := d.SessionManager.Load(r)
	b, err := session.GetBytes("json")
	if err != nil {
		return
	}
	if len(b) == 0 {
		err = fmt.Errorf("Session not authenticated")
		return
	}

	err = json.Unmarshal(b, &u)
	return
}

// Data Interfaces

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
		// No match, UserId not valid
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
