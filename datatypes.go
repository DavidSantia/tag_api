package tag_api

import (
	"github.com/alexedwards/scs"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func NewData() (data *ApiData) {
	data = &ApiData{}
	data.InitSessions()
	d = data
	return
}

// Local data - most functions are methods of this
type ApiData struct {
	Debug          bool
	Logfile        string
	Router         *httprouter.Router
	Db             *sqlx.DB
	UserMap        UserMap
	GroupMap       GroupMap
	ImageMap       ImageMap
	SessionManager *scs.Manager
}

type UserMap map[int64]User

type GroupMap map[int64]Group

type ImageMap map[int64]Image

type ImagesGroupsMap map[int64]bool

type JwtPayload struct {
	UserId int64  `json:"id"`
	Guid   string `json:"guid"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
