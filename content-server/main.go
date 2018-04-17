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

func main() {

	data := tag_api.NewData()
	err := data.GetCmdLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize log
	var level tag_api.Level = tag_api.LogINFO
	if data.Debug {
		level = tag_api.LogDEBUG
	}
	tag_api.NewLog(level, data.Logfile)

	name := "Content"
	tag_api.Log.Info.Printf("-------- %s API Server [Version %s-%s Build %s %s] --------",
		name, GitBranch, GitCommit, GitState, BuildDate)

	// Initialize HTTP router
	data.Router = tag_api.NewContentRouter()

	// Connect SQL DB
	err = data.ConnectDB()
	if err != nil {
		tag_api.Log.Error.Println(err)
		os.Exit(1)
	}
	defer data.Db.Close()

	// Connect Bolt DB
	err = data.ConnectBolt()
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

	data.StartServer(":8080", name)
	os.Exit(0)
}
