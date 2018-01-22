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

	name := "Authenticate"
	tag_api.Log.Info.Printf("-------- %s API Server [Version %s-%s Build %s %s] --------",
		name, GitBranch, GitCommit, GitState, BuildDate)

	// Initialize HTTP router
	data.Router = tag_api.NewAuthRouter()

	// Connect SQL DB
	err = data.ConnectDB()
	if err != nil {
		tag_api.Log.Error.Println(err)
		os.Exit(1)
	}

	// Load users
	data.LoadUsers()

	data.StartServer(":8081", name)
	os.Exit(0)
}
