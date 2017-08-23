package tag_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/engine/memstore"
	"github.com/alexedwards/scs/session"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/julienschmidt/httprouter"
)

func (data *ApiData) InitSessions() {

	// Initialise a new storage engine
	engine := memstore.New(0)

	// Initialise the session manager middleware
	data.SessionManager = session.Manage(engine,
		// maximum length of time a session can be inactive
		session.IdleTimeout(30*time.Minute),
		// maximum length of time that a session is valid
		session.Lifetime(8*time.Hour),
		// whether the session cookie should be retained after a user closes their browser
		session.Persist(false),
	)
}

// HTTP Handlers

func HandleAuthenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u User
	var b []byte

	auth := r.Header.Get("Authorization")
	if len(auth) == 0 || !strings.Contains(auth, "Bearer") {
		HandleError(w, http.StatusBadRequest, "Invalid authentication: "+auth)
		return
	}

	// Decode payload from JSON Web Token
	token := strings.TrimSpace(strings.Replace(auth, "Bearer", "", 1))
	payload, headers, err := jose.Decode(token, JwtKey)
	if err != nil {
		HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	logging.Debug.Printf("Headers = %+v\n", headers)
	logging.Debug.Printf("Payload = %s\n", payload)

	// Make sure payload is JSON
	if len(payload) == 0 || payload[0] != '{' {
		HandleError(w, http.StatusBadRequest, "Payload decryption failed")
		return
	}

	logging.Debug.Printf("Authenticate Payload: %s\n", payload)

	pl := JwtPayload{}
	err = json.Unmarshal([]byte(payload), &pl)
	if err != nil {
		HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Lookup user by id
	u, err = d.UserFind(pl)
	if err != nil {
		HandleError(w, http.StatusBadRequest, err.Error())
		return
	}
	b, err = json.Marshal(u)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Store member data in session
	gid := fmt.Sprintf("%d", u.GroupId)
	err = session.PutString(r, "gid", gid)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = session.PutBytes(r, "json", b)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
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
	gid, err = session.GetString(r, "gid")
	if len(gid) == 0 {
		err = fmt.Errorf("Session not authenticated")
	}
	group_id, err = strconv.ParseInt(gid, 10, 64)
	return
}

func GetUserFromSession(r *http.Request) (u User, err error) {

	// See if session is authenticated
	b, err := session.GetBytes(r, "json")
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

func (data *ApiData) UserFind(pl JwtPayload) (u User, err error) {
	var query string
	u = User{}

	// Validate payload
	if pl.UserId == 0 {
		err = fmt.Errorf("Invalid UserId in payload")
		return
	}
	if len(pl.Guid) == 0 {
		err = fmt.Errorf("Invalid Guid in payload")
		return
	}

	// Lookup member
	query = data.MakeQuery(u, UserQuery, pl.UserId)
	logging.Debug.Printf("Query: %s\n", query)

	err = data.Db.QueryRowx(query).StructScan(&u)
	if err == sql.ErrNoRows {
		// No match, UserId not valid
		err = fmt.Errorf("UserId %d not valid", pl.UserId)
	}
	if err != nil {
		return
	}

	// Validate Guid
	if u.Guid != pl.Guid {
		err = fmt.Errorf("Invalid Guid %s specified for User %d", pl.Guid, pl.UserId)
		return
	}

	// Validate Status
	if !u.Status {
		err = fmt.Errorf("Expired User %d", pl.UserId)
	}
	return
}
