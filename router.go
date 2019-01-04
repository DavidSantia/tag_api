package tag_api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() (router *httprouter.Router) {

	router = httprouter.New()

	router.Handle("GET", "/", handleIndex)
	router.Handle("GET", "/authenticate", HandleAuthTester)
	router.Handle("POST", "/authenticate", HandleAuthenticate)
	router.Handle("GET", "/keepalive", HandleAuthKeepAlive)
	router.Handle("GET", "/image", HandleAllImages)
	router.Handle("GET", "/image/:Id", HandleImage)
	router.Handle("GET", "/user", HandleUser)

	return
}

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<b>API Demo endpoints</b>")
	fmt.Fprint(w, "<ul><li>GET /authenticate</li>")
	fmt.Fprint(w, "<li>POST /authenticate</li>")
	fmt.Fprint(w, "<li>GET /keepalive</li></ul>")
	fmt.Fprint(w, "<ul><li>GET /image</li>")
	fmt.Fprint(w, "<li>GET /image/Id</li>")
	fmt.Fprint(w, "<li>GET /user</li></ul>")
}

func (data *ApiData) StartServer(hostApi, portApi, hostNATS, portNATS string) {
	var err error

	err = data.ConnectNATS(hostNATS, portNATS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer data.NConn.Close()

	go data.ListenNATSSub()

	apiUrl := hostApi + ":" + portApi

	Log.Info.Println("API Ready on " + apiUrl)
	err = http.ListenAndServe(apiUrl, data.SessionManager.Use(data.Router))
	fmt.Println(err)
	os.Exit(1)
}
