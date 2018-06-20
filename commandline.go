package tag_api

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Validate command-line arguments
func (data *ApiData) validateFlags() (err error) {
	var absolutePath, relPath string
	var dirInfo os.FileInfo

	// Validate Logfile
	if len(data.Logfile) != 0 {
		relPath = filepath.Dir(data.Logfile)
		if absolutePath, err = filepath.Abs(data.Logfile); err != nil {
			return fmt.Errorf("-logfile %s %v", data.Logfile, err)
		}
		dirInfo, err = os.Stat(relPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("-logfile %s: directory %s %v", data.Logfile, relPath, err)
		} else if !dirInfo.IsDir() {
			return fmt.Errorf("-logfile %s: %s is not a directory", data.Logfile, relPath)
		}
		data.Logfile = absolutePath
	}
	return
}

func (data *ApiData) GetCmdLine() (err error) {

	// Define command-line arguments
	flag.StringVar(&data.Logfile, "log", "", "Specify logging filename")
	flag.StringVar(&data.DbHost, "dbhost", DbHost, "Specify DB host")
	flag.StringVar(&data.DbPort, "dbport", DbPort, "Specify DB port")
	flag.BoolVar(&data.Debug, "debug", false, "Debug logging")
	flag.StringVar(&data.NHost, "nhost", NHost, "Specify NATS host")

	// Parse commandline flag arguments
	flag.Parse()

	// Validate
	err = data.validateFlags()
	return
}
