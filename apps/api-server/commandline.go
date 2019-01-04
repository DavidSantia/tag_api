package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Settings struct {
	debug    bool
	logfile  string
	server   string
	hostDb   string
	portDb   string
	nameDb   string
	hostNATS string
	portNATS string
	hostApi  string
	portApi  string
}

// Validate command-line arguments
func (settings *Settings) validateFlags() (err error) {
	var absolutePath, relPath string
	var dirInfo os.FileInfo

	// Validate Logfile
	if len(settings.logfile) != 0 {
		relPath = filepath.Dir(settings.logfile)
		if absolutePath, err = filepath.Abs(settings.logfile); err != nil {
			return fmt.Errorf("-logfile %s %v", settings.logfile, err)
		}
		dirInfo, err = os.Stat(relPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("-logfile %s: directory %s %v", settings.logfile, relPath, err)
		} else if !dirInfo.IsDir() {
			return fmt.Errorf("-logfile %s: %s is not a directory", settings.logfile, relPath)
		}
		settings.logfile = absolutePath
	}
	return
}

func (settings *Settings) getCmdLine() (err error) {

	// Define command-line arguments
	flag.BoolVar(&settings.debug, "debug", false, "Debug logging")
	flag.StringVar(&settings.logfile, "log", "", "Specify logging filename")
	flag.StringVar(&settings.hostApi, "host", "", "Specify Api host")
	flag.StringVar(&settings.portApi, "port", "8080", "Specify Api port")
	flag.StringVar(&settings.hostDb, "dbhost", "", "Specify DB host")
	flag.StringVar(&settings.portDb, "dbport", "3306", "Specify DB port")
	flag.StringVar(&settings.hostNATS, "nhost", "", "Specify NATS host")
	flag.StringVar(&settings.portNATS, "nport", "4222", "Specify NATS port")

	// Parse commandline flag arguments
	flag.Parse()

	// Validate
	err = settings.validateFlags()
	return
}
