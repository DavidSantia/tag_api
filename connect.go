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

func (data *ApiData) ConnectDB(user, pass, name, host, port string) (err error) {

	// Set DB connection resource string
	resource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	Log.Info.Printf("Connecting to %s on %s", name, host)
	// Retry connection if DB still initializing
	for i := 0; i < DbConnectRetries; i++ {
		data.Db, err = sqlx.Connect("mysql", resource)
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

func (data *ApiData) ConnectBolt(file string) (err error) {

	Log.Info.Printf("Connecting to %s", file)
	data.BoltDb, err = bolt.Open(file, 0644, nil)

	// Bucket name
	data.BoltBucket = []byte("Content")
	return
}

func (data *ApiData) ConnectNATS(host, port string) (err error) {

	natsUrl := "nats://" + host + ":" + port

	Log.Info.Printf("Connecting to %s", natsUrl)
	data.NConn, err = nats.Connect(natsUrl)
	return
}
