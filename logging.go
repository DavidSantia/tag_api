package tag_api

import (
	"io/ioutil"
	"log"
	"os"
)

type Logging struct {
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Debug *log.Logger
}

type Level uint

const (
	LogNONE Level = iota
	LogERROR
	LogWARN
	LogINFO
	LogDEBUG
)

var Log Logging

func NewLog(level Level, logfile string) {
	var fLog *os.File = os.Stderr
	var err error

	fError := ioutil.Discard
	fWarn := ioutil.Discard
	fInfo := ioutil.Discard
	fDebug := ioutil.Discard

	// Output to logfile, if specified
	if len(logfile) != 0 {
		fLog, err = os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("Error opening logfile: %v\n", err)
			fLog = os.Stderr
		}
	}

	if level > LogNONE {
		fError = fLog
	}
	if level > LogERROR {
		fWarn = fLog
	}
	if level > LogWARN {
		fInfo = fLog
	}
	if level > LogINFO {
		fDebug = fLog
	}

	Log = Logging{
		Error: log.New(fError, "ERROR: ", log.LstdFlags),
		Warn:  log.New(fWarn, " WARN: ", log.LstdFlags),
		Info:  log.New(fInfo, " INFO: ", log.LstdFlags),
		Debug: log.New(fDebug, "DEBUG: ", log.LstdFlags),
	}
}
