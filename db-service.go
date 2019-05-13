package tag_api

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/newrelic/go-agent"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DbService struct {
	settings DbSettings
	db       *sqlx.DB
}

type DbSettings struct {
	userDb string
	passDb string
	nameDb string
	hostDb string
	portDb string
}

func NewDbService(user, pass, name, host, port string) *DbService {
	// Configure Db settings
	ds := DbService{}
	ds.settings.userDb = user
	ds.settings.passDb = pass
	ds.settings.nameDb = name
	ds.settings.hostDb = host
	ds.settings.portDb = port
	return &ds
}

func (ds *DbService) Connect() (err error) {

	// Set DB connection resource string
	resource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", ds.settings.userDb, ds.settings.passDb,
		ds.settings.hostDb, ds.settings.portDb, ds.settings.nameDb)

	Log.Info.Printf("Connecting to %s on %s", ds.settings.nameDb, ds.settings.hostDb)
	// Retry connection if DB still initializing
	for i := 0; i < DbConnectRetries; i++ {
		ds.db, err = sqlx.Connect("mysql", resource)
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

func (ds *DbService) Close() {
	ds.db.Close()
}

func (ds *DbService) Queryx(query string) (rows *sqlx.Rows, err error) {
	var retry int

	s := newrelic.DatastoreSegment{
		Product:            newrelic.DatastoreMySQL,
		Operation:          "SELECT",
		ParameterizedQuery: query,
		Host:               ds.settings.hostDb,
		PortPathOrID:       ds.settings.portDb,
		DatabaseName:       ds.settings.nameDb,
	}
	s.StartTime = newrelic.StartSegmentNow(CurrentTxn)
	defer s.End()

	for retry = 0; retry <= DbConnectRetries; retry++ {
		rows, err = ds.db.Queryx(query)
		if err != nil {
			if err == mysql.ErrInvalidConn {
				Log.Warn.Printf("retry #%d MySQL Invalid Connection", retry+1)
				continue
			}
			return
		}
		break
	}
	return
}

// Helper functions

func makeQuery(dt interface{}, fromClause string, v ...interface{}) (query string) {
	var tag, sqlTag string
	var ok bool
	var i int
	var params []string

	st := reflect.TypeOf(dt)

	// Iterate through struct fields, gathering db tags
	for i = 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if tag, ok = field.Tag.Lookup("db"); ok {
			if len(tag) > 0 {
				// If sql tag specified, override db tag
				sqlTag, ok = field.Tag.Lookup("sql")
				if ok {
					if sqlTag == "-" {
						// Disable generate (i.e., SELECT param included in fromClause)
						continue
					}
					tag = sqlTag + " AS " + tag
				}
				params = append(params, tag)
			}
		}
	}

	// Generate SELECT paramas from struct tags, then add FROM clause
	query = "SELECT " + strings.Join(params, ", ") +
		strings.Replace(fmt.Sprintf(fromClause, v...), "FROM", "\nFROM", 1)
	return
}
