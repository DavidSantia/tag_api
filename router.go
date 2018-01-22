package tag_api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
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

	return
}

func AuthIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<b>API Demo authenticate endpoints</b>")
	fmt.Fprint(w, "<ul><li>GET /authenticate</li>")
	fmt.Fprint(w, "<li>POST /authenticate</li></ul>")
}

func (data *ApiData) StartServer(host, name string) {
	Log.Info.Println(name, "API Ready")

	err := http.ListenAndServe(host, data.Router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
