package util

import (
	"fmt"
	. "os"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func InitializeLogging(logFile string) {
	var file, err = OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0644)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	Log().Info("l'oggo Init!")
}

func Log() *log.Entry {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	f := frame.Function
	if idx := strings.LastIndex(frame.Function, "/"); idx >= 0 {
		f = f[idx+1:]
	}

	return log.WithField("timestamp", time.Now().Local().Format(time.RFC3339)).
		WithField("func", f).
		WithField("line", frame.Line)
}
