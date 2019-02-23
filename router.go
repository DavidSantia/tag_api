package tag_api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/newrelic/go-agent"
)

func (data *ApiData) NewRouter(cs ContentService, app newrelic.Application) (router *httprouter.Router) {

	data.router = httprouter.New()
	data.router.Handle("GET", "/", WrapRouterHandle(app, handleIndex))
	data.router.Handle("GET", "/authenticate", WrapRouterHandle(app, handleAuthTestpage))
	data.router.Handle("POST", "/authenticate", WrapRouterHandle(app, makeHandleAuthenticate(cs)))
	data.router.Handle("GET", "/keepalive", WrapRouterHandle(app, makeHandleAuthKeepAlive(cs)))
	data.router.Handle("GET", "/image", WrapRouterHandle(app, makeHandleAllImages(cs)))
	data.router.Handle("GET", "/image/:Id", WrapRouterHandle(app, makeHandleImage(cs)))
	data.router.Handle("GET", "/user", WrapRouterHandle(app, makeHandleUser(cs)))

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

func (data *ApiData) StartServer() {

	Log.Info.Println("API Ready on " + data.apiUrl)
	err := http.ListenAndServe(data.apiUrl, data.sessionManager.Use(data.router))
	fmt.Println(err)
	os.Exit(1)
	return
}
