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

	tag_api.Log.Info.Printf("-------- API Server [Version %s-%s Build %s %s] --------",
		GitBranch, GitCommit, GitState, BuildDate)

	// Initialize HTTP router
	data.Router = tag_api.NewRouter()

	// Connect SQL DB
	err = data.ConnectDB()
	if err != nil {
		tag_api.Log.Error.Println(err)
		os.Exit(1)
	}

	data.AutoLoad()

	data.StartServer()
	os.Exit(0)
}
