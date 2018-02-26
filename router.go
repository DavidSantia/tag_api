package tag_api

import (
	"fmt"
	"net/http"
	"os"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/go-nats"
)

func NewContentRouter() (router *httprouter.Router) {

	router = httprouter.New()

	router.Handle("GET", "/", ContentIndex)
	router.Handle("GET", "/image", HandleAllImages)
	router.Handle("GET", "/image/:Id", HandleImage)
	router.Handle("GET", "/user", HandleUser)

	return
}

func ContentIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<b>API Demo Content endpoints</b>")
	fmt.Fprint(w, "<ul><li>GET /image</li>")
	fmt.Fprint(w, "<li>GET /image/Id</li>")
	fmt.Fprint(w, "<li>GET /user</li></ul>")
}

func NewAuthRouter() (router *httprouter.Router) {

	router = httprouter.New()

	router.Handle("GET", "/", AuthIndex)
	router.Handle("GET", "/authenticate", HandleAuthTester)
	router.Handle("POST", "/authenticate", HandleAuthenticate)
	router.Handle("GET", "/keepalive", HandleAuthKeepAlive)

	return
}

func AuthIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<b>API Demo authenticate endpoints</b>")
	fmt.Fprint(w, "<ul><li>GET /authenticate</li>")
	fmt.Fprint(w, "<li>POST /authenticate</li>")
	fmt.Fprint(w, "<li>GET /keepalive</li></ul>")
}

func (data *ApiData) ListenUserChan() {
	var qMsg QueueMessage
	ch := make(chan *nats.Msg, 64)
	q := "users"

	sub, err := data.Nconn.ChanSubscribe(q, ch)
	if err != nil {
		Log.Error.Println(err)
		return
	}
	defer sub.Unsubscribe()

	Log.Info.Printf("Subscribed to %q on %s\n", q, NATSUrl)
	for {
		msg := <-ch
		err = json.Unmarshal(msg.Data, &qMsg)
		if err != nil {
			Log.Error.Println(err)
		}

		switch qMsg.Command {
		case "adduser":
			err = data.AddUser(msg.Data)
			if err != nil {
				Log.Error.Printf("Add User: %v\n", err)
			}
		default:
			Log.Info.Printf("Unrecognized command: %s\n", qMsg.Command)
		}
	}
}

func (data *ApiData) StartServer(host, name string) {
	var err error

	Log.Info.Println(name, "API Ready")

	data.Nconn, err = nats.Connect(NATSUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if name == "Content" {
		go data.ListenUserChan()
	}

	err = http.ListenAndServe(host, data.SessionManager.Use(data.Router))
	data.Nconn.Close()
	fmt.Println(err)
	os.Exit(1)
}
