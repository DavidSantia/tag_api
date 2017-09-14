package tag_api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() (router *httprouter.Router) {

	router = httprouter.New()

	router.Handle("GET", "/", Index)
	router.Handle("GET", "/authenticate", HandleAuthTester)
	router.Handle("POST", "/authenticate", HandleAuthenticate)
	router.Handle("GET", "/image", HandleAllImages)
	router.Handle("GET", "/image/:Id", HandleImage)
	router.Handle("GET", "/user", HandleUser)

	return
}

func (data *ApiData) StartServer() {
	Log.Info.Println("API Ready")

	err := http.ListenAndServe(":8080", data.Router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
