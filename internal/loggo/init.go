package loggo

import (
	"os"
	"path"

	"github.com/egor3f/loggo/internal/util"
)

const (
	parentPath = ".loggo"
	logsPath   = "logs"
	currentLog = "latest.log"
)

var LatestLog string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	//now := time.Now().Local().Format("2006.01.02T15.04.05")
	//file := fmt.Sprintf("%s.log", now)
	paramsDir := path.Join(home, parentPath, logsPath)

	if err := os.MkdirAll(paramsDir, os.ModePerm); err != nil {
		panic(err)
	}
	//prev := path.Join(paramsDir, file)
	LatestLog = path.Join(paramsDir, currentLog)
	//os.Rename(LatestLog, prev)

	util.InitializeLogging(LatestLog)
}
