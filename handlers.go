package tag_api

import (
	"encoding/json"
	"fmt"
	"github.com/newrelic/go-agent"
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

var CurrentTxn newrelic.Transaction

func WrapRouterHandle(app newrelic.Application, handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		Log.Debug.Printf("%s %s", r.Method, r.RequestURI)
		if app != nil {
			CurrentTxn = app.StartTransaction(r.RequestURI, w, r)
			defer CurrentTxn.End()

			r = newrelic.RequestWithTransactionContext(r, CurrentTxn)
		}

		handle(w, r, ps)
	}
}
