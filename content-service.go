package tag_api

import (
	"sync"

	"github.com/boltdb/bolt"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/go-nats"
)

type ContentService interface {
	ConfigureDB(user, pass, name, host, port string)
	CloseDB()
	ConfigureNATS(host, port, channel string)
	ConnectNATS() (err error)
	CloseNATS()
	EnableLoadAll()
	GetUser(id int64) (user User, ok bool)
	GetGroup(id int64) (group Group, ok bool)
	GetImage(id int64) (image Image, ok bool)
	ListenForUpdates()
	LoadCacheUpdates() (err error)
	LoadFromDb() (err error)
	PublishUpdate() (err error)
	ShowUpdates()
	StoreDbUpdates()
	UpdateFromCache()
}

func NewContentService(boltFile, boltBucket string) (cs ContentService) {
	// Configure BoltDb settings
	bs := BoltService{mutex: &sync.RWMutex{}}
	bs.settings.boltFile = boltFile
	bs.settings.boltBucket = []byte(boltBucket)
	return &bs
}

type BoltService struct {
	settings    Settings
	UserMap     UserMap
	GroupMap    GroupMap
	ImageMap    ImageMap
	refresh     bool
	updateGroup []byte
	updateImage []byte
	updateUser  []byte
	boltDb      *bolt.DB
	db          *sqlx.DB
	nconn       *nats.Conn
	mutex       *sync.RWMutex
}

type Settings struct {
	enableGroups       bool
	enableImages       bool
	enableImagesGroups bool
	enableUsers        bool
	boltBucket         []byte
	boltFile           string
	userDb             string
	passDb             string
	nameDb             string
	hostDb             string
	portDb             string
	hostNATS           string
	portNATS           string
	channelNATS        string
}

func (bs *BoltService) ConfigureDB(user, pass, name, host, port string) {
	// Configure Db settings
	bs.settings.userDb = user
	bs.settings.passDb = pass
	bs.settings.nameDb = name
	bs.settings.hostDb = host
	bs.settings.portDb = port
}

func (bs *BoltService) CloseDB() {
	bs.db.Close()
}

func (bs *BoltService) ConfigureNATS(host, port, channel string) {
	// Configure NATS settings
	bs.settings.hostNATS = host
	bs.settings.portNATS = port
	bs.settings.channelNATS = channel
}

func (bs *BoltService) CloseNATS() {
	bs.nconn.Close()
}

func (bs *BoltService) EnableLoadAll() {
	bs.settings.enableGroups = true
	bs.settings.enableImages = true
	bs.settings.enableImagesGroups = true
	bs.settings.enableUsers = true
}

func (bs *BoltService) GetUser(id int64) (user User, ok bool) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	user, ok = bs.UserMap[id]
	return
}

func (bs *BoltService) GetGroup(id int64) (group Group, ok bool) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	group, ok = bs.GroupMap[id]
	return
}

func (bs *BoltService) GetImage(id int64) (image Image, ok bool) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	image, ok = bs.ImageMap[id]
	return
}

func (bs *BoltService) ListenForUpdates() {

}

func (bs *BoltService) LoadCacheUpdates() (err error) {
	return
}

func (bs *BoltService) LoadFromDb() (err error) {
	// Connect Db if not intialized
	if bs.db == nil {
		if err = bs.connectDB(); err != nil {
			return
		}
	}
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	// Load groups
	if bs.settings.enableGroups {
		bs.loadGroups()
	}

	// Load images
	if bs.settings.enableImages {
		bs.loadImages()
	}

	// Load map of images for each group
	if bs.settings.enableImagesGroups {
		bs.loadImagesGroups()
	}

	// Load users
	if bs.settings.enableUsers {
		bs.loadUsers()
	}
	return
}

func (bs *BoltService) PublishUpdate() (err error) {
	return
}

func (bs *BoltService) ShowUpdates() {

}

func (bs *BoltService) StoreDbUpdates() {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	// Store groups
	if bs.settings.enableGroups {
		bs.storeGroups()
	}

	// Store images
	if bs.settings.enableImages {
		bs.storeImages()
	}

	// Store users
	if bs.settings.enableUsers {
		bs.storeUsers()
	}
	return
}

func (bs *BoltService) UpdateFromCache() {

}
