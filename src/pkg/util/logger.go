package util

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func WriteLog() {
	Log.Out = os.Stdout

	currentTime := time.Now()
	// set format YYYY-MM-DD for rotate log file
	now := currentTime.Format("2006-01-02")

	file, err := os.OpenFile("./logs/log-"+now+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stder")
	}
}
