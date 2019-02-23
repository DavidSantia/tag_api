package main

import (
	"fmt"
	"os"

	"github.com/DavidSantia/tag_api"
	"github.com/newrelic/go-agent"
)

// These fields are populated by govvv
var (
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
)

const (
	// Default Bolt DB file
	BoltDB = "./content.db"

	// Retries to wait for docker DB instance
	DbConnectRetries = 5

	// MySQL DB info
	DbUser = "demo"
	DbPass = "welcome1"
	DbName = "tagdemo"

	// NATS server
	NSub = "update"
)

func main() {
	var app newrelic.Application

	settings := Settings{server: "Tag Api"}

	err := settings.getCmdLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := tag_api.NewData(settings.hostApi, settings.portApi)

	// Initialize log
	var level tag_api.Level = tag_api.LogINFO
	if settings.debug {
		level = tag_api.LogDEBUG
	}
	tag_api.NewLog(level, settings.logFile)

	tag_api.Log.Info.Printf("-------- %s Server [Version %s-%s Build %s %s] --------",
		settings.server, GitBranch, GitCommit, GitState, BuildDate)

	if len(settings.apmKey) > 0 {
		// Initialize New Relic agent
		tag_api.Log.Info.Println("Initializing APM")
		config := newrelic.NewConfig(settings.server, settings.apmKey)
		app, err = newrelic.NewApplication(config)
		if err != nil {
			tag_api.Log.Error.Printf("Error initializing APM: %v", err)
			return
		}
	}

	// Initialize content service
	cs := tag_api.NewContentService(settings.boltFile, DbName)

	// This illustrates two ways of using the content service
	if settings.loadDb {

		// Option 1: Load all content from Db
		ds := tag_api.NewDbService(DbUser, DbPass, DbName, settings.hostDb, settings.portDb)
		err = ds.Connect()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer ds.Close()
		cs.ConfigureDbService(ds)

		cs.EnableLoadAll()
		err = cs.LoadFromDb()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {

		// Option 2: Listen to NATS messaging for content updates from the updater service
		cs.ConfigureNATS(settings.hostNATS, settings.portNATS, NSub)

		err = cs.ConnectNATS()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer cs.CloseNATS()
	}

	// Initialize HTTP router
	data.NewRouter(cs, app)

	data.StartServer()
	os.Exit(0)
}
