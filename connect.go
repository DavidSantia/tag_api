package tag_api

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func (bs *BoltService) connectDB() (err error) {

	// Set DB connection resource string
	resource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", bs.settings.userDb, bs.settings.passDb,
		bs.settings.hostDb, bs.settings.portDb, bs.settings.nameDb)

	Log.Info.Printf("Connecting to %s on %s", bs.settings.nameDb, bs.settings.hostDb)
	// Retry connection if DB still initializing
	for i := 0; i < DbConnectRetries; i++ {
		bs.db, err = sqlx.Connect("mysql", resource)
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				time.Sleep(10 * time.Second)
				Log.Info.Printf("Retry connection #%d", i+1)
				continue
			}
			return
		}
	}
	return
}

func (bs *BoltService) connectBolt() (err error) {

	Log.Info.Printf("Connecting to %s", bs.settings.boltFile)
	bs.boltDb, err = bolt.Open(bs.settings.boltFile, 0644, nil)
	return
}

func (bs *BoltService) ConnectNATS() (err error) {

	natsUrl := "nats://" + bs.settings.hostNATS + ":" + bs.settings.portNATS

	Log.Info.Printf("Connecting to %s", natsUrl)
	bs.nconn, err = nats.Connect(natsUrl)
	return
}
