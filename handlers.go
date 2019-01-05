package tag_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, "<b>API Demo using Go-lang struct tags</b>")
}

func HandleReply(w http.ResponseWriter, status int, j string) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	fmt.Fprintln(w, j)
}

func HandleError(w http.ResponseWriter, status int, uri string, err error) {
	Log.Error.Printf("%q %v", uri, err)

	w.WriteHeader(status)
	rMsg := ResponseMessage{
		Message: fmt.Sprintf("Error: %v", err),
		Status:  http.StatusText(status),
	}
	b, _ := json.Marshal(rMsg)
	fmt.Fprintln(w, string(b))
}
