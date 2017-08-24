package tag_api

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func (data *ApiData) ConnectDB() (err error) {
	var resource, host, port, user, pass, name string

	// Set DB connection resource string
	host = DbHost
	port = DbPort
	user = DbUser
	pass = DbPass
	name = DbName
	resource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	Log.Info.Printf("Connecting to %s on %s", name, host)
	data.Db, err = sqlx.Connect("mysql", resource)
	if err != nil {
		return err
	}
	return
}
