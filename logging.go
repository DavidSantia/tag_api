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
	logNONE Level = iota
	logERROR
	logWARN
	logINFO
	logDEBUG
)

var logging Logging

func NewLog(level Level) {
	fError := ioutil.Discard
	fWarn := ioutil.Discard
	fInfo := ioutil.Discard
	fDebug := ioutil.Discard

	if level > logNONE {
		fError = os.Stderr
	}
	if level > logERROR {
		fWarn = os.Stderr
	}
	if level > logWARN {
		fInfo = os.Stderr
	}
	if level > logINFO {
		fDebug = os.Stderr
	}

	logging = Logging{
		Error: log.New(fError, "ERROR: ", log.LstdFlags),
		Warn:  log.New(fWarn, " WARN: ", log.LstdFlags),
		Info:  log.New(fInfo, " INFO: ", log.LstdFlags),
		Debug: log.New(fDebug, "DEBUG: ", log.LstdFlags),
	}
}
