package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Settings struct {
	debug    bool
	loadDb   bool
	apmKey   string
	logFile  string
	server   string
	boltFile string
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
	if len(settings.logFile) > 0 {
		relPath = filepath.Dir(settings.logFile)
		if absolutePath, err = filepath.Abs(settings.logFile); err != nil {
			return fmt.Errorf("-log %s %v", settings.logFile, err)
		}
		dirInfo, err = os.Stat(relPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("-log %s: directory %s %v", settings.logFile, relPath, err)
		} else if !dirInfo.IsDir() {
			return fmt.Errorf("-log %s: %s is not a directory", settings.logFile, relPath)
		}
		settings.logFile = absolutePath
	}

	// Validate BoltDb file
	if len(settings.logFile) > 0 {
		relPath = filepath.Dir(settings.logFile)
		if absolutePath, err = filepath.Abs(settings.logFile); err != nil {
			return fmt.Errorf("-logfile %s %v", settings.logFile, err)
		}
		dirInfo, err = os.Stat(relPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("-logfile %s: directory %s %v", settings.logFile, relPath, err)
		} else if !dirInfo.IsDir() {
			return fmt.Errorf("-logfile %s: %s is not a directory", settings.logFile, relPath)
		}
		settings.logFile = absolutePath
	}

	// Validate APM key
	if len(settings.apmKey) > 0 {
		if len(settings.apmKey) != 40 {
			return fmt.Errorf("-apmkey: must be 40 characters")
		}
	}

	return
}

func (settings *Settings) getCmdLine() (err error) {

	// Define command-line arguments
	flag.BoolVar(&settings.debug, "debug", false, "Debug logging")
	flag.BoolVar(&settings.loadDb, "dbload", false, "Load from DB instead of BoltDB")
	flag.StringVar(&settings.apmKey, "apmkey", "", "Specify APM license key")
	flag.StringVar(&settings.logFile, "log", "", "Specify logging filename")
	flag.StringVar(&settings.boltFile, "bolt", BoltDB, "Specify BoltDB filename")
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
