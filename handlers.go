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

var d *ApiData

func HandleError(w http.ResponseWriter, status int, message string) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	status_txt := http.StatusText(status)
	logging.Error.Printf("%s: %s\n", status_txt, message)
	eResp := ErrorResponse{
		Error:  message,
		Status: status_txt,
	}
	b, _ := json.Marshal(eResp)
	fmt.Fprintln(w, string(b))
}

func HandleStatus(w http.ResponseWriter, status int) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	status_txt := http.StatusText(status)
	sResp := StatusResponse{
		Status: status_txt,
	}
	b, _ := json.Marshal(sResp)
	fmt.Fprintln(w, string(b))
}

func HandleReply(w http.ResponseWriter, status int, j string) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	fmt.Fprintln(w, j)
}
