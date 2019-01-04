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
	// Bolt DB file
	BoltDB = "./content.db"

	// Retries to wait for docker DB instance
	DbConnectRetries = 5

	// MySQL DB info
	DbUser = "demo"
	DbPass = "welcome1"
	DbName = "tagdemo"

	// NATS server
	NHost = "localhost"
	NSub  = "update"
)


func main() {

	data := tag_api.NewData()
	settings := Settings{server: "Api"}

	err := settings.getCmdLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize log
	var level tag_api.Level = tag_api.LogINFO
	if settings.debug {
		level = tag_api.LogDEBUG
	}
	tag_api.NewLog(level, settings.logfile)

	tag_api.Log.Info.Printf("-------- %s Server [Version %s-%s Build %s %s] --------",
		settings.server, GitBranch, GitCommit, GitState, BuildDate)

	// Initialize HTTP router
	data.Router = tag_api.NewRouter()

	// Connect SQL DB
	err = data.ConnectDB(DbUser, DbPass, DbName, settings.hostDb, settings.portDb)
	if err != nil {
		tag_api.Log.Error.Println(err)
		os.Exit(1)
	}
	defer data.Db.Close()

	// Connect Bolt DB
	err = data.ConnectBolt(BoltDB)
	if err != nil {
		tag_api.Log.Error.Println(err)
		os.Exit(1)
	}
	defer data.BoltDb.Close()

	// Load images
	data.LoadImages()

	// Store images
	data.StoreImages()

	// Load groups
	data.LoadGroups()

	// Load map of images for each group
	data.LoadImagesGroups()

	// Store groups
	data.StoreGroups()

	// Refresh groups and images
	data.RefreshImages()
	data.RefreshGroups()

	data.StartServer(settings.hostApi, settings.portApi, settings.hostNATS, settings.portNATS)
	os.Exit(0)
}
