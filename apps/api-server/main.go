package main

import (
	"fmt"
	"os"

	"github.com/DavidSantia/tag_api"
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

	settings := Settings{server: "Api"}

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

	// Initialize content service
	cs := tag_api.NewContentService(settings.boltFile, DbName)

	// Configure Db
	cs.ConfigureDB(DbUser, DbPass, DbName, settings.hostDb, settings.portDb)
	defer cs.CloseDB()

	// Configure NATS
	cs.ConfigureNATS(settings.hostNATS, settings.portNATS, NSub)
	err = cs.ConnectNATS()
	if err != nil {
		return
	}
	defer cs.CloseNATS()

	// Load content from Db
	cs.EnableLoadAll()
	cs.LoadFromDb()

	// Initialize HTTP router
	data.NewRouter(cs)

	data.StartServer()
	os.Exit(0)
}
