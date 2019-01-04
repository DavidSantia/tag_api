package tag_api

import (
	"github.com/alexedwards/scs"
	"github.com/boltdb/bolt"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/go-nats"
)

func NewData() (data *ApiData) {
	data = &ApiData{
		UserMap:  make(UserMap),
		ImageMap: make(ImageMap),
		GroupMap: make(GroupMap),
	}
	data.InitSessions()
	d = data
	return
}

// Local data - most functions are methods of this
type ApiData struct {
	BoltDb         *bolt.DB
	BoltBucket     []byte
	Db             *sqlx.DB
	NConn          *nats.Conn
	Router         *httprouter.Router
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

type ResponseMessage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type QueueMessage struct {
	Command string `json:"command"`
}
