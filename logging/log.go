package logging

import (
	"io"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logger() {
	if os.Getenv("DEBUG") == "true" {
		logFileName := "collaborative_story_" + time.Now().Format("20060102150405") + ".log"

		logFile, _ := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o644)
		Formatter := new(log.TextFormatter)
		Formatter.FullTimestamp = true
		log.SetFormatter(Formatter)
		log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	} else {
		log.SetOutput(ioutil.Discard)
	}
	log.Infoln("DEBUG =", os.Getenv("DEBUG"))
}
