package tag_api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HTTP Handlers

func HandleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var u User

	u, err := GetUserFromSession(r)
	if err != nil {
		HandleError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var b []byte
	b, err = json.Marshal(u)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	HandleReply(w, http.StatusOK, string(b))
}
