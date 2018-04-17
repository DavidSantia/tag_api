package tag_api

import (
	"fmt"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func (data *ApiData) ConnectDB() (err error) {
	var resource, user, pass, name string

	// Set DB connection resource string
	user = DbUser
	pass = DbPass
	name = DbName
	resource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, data.DbHost, data.DbPort, name)

	Log.Info.Printf("Connecting to %s on %s", name, data.DbHost)
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

func (data *ApiData) ConnectBolt() (err error) {
	var file string = BoltDB

	Log.Info.Printf("Connecting to %s", file)
	data.BoltDb, err = bolt.Open(file, 0644, nil)

	// Bucket name
	data.BoltBucket = []byte("Content")

	return
}
