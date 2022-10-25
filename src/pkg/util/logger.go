package util

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func MakeLogger(path string, display bool) *logrus.Logger {
	currentTime := time.Now()
	now := currentTime.Format("2006-01-02")

	f, err := os.OpenFile("./logs/"+path+"/log-"+now+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}

	logger := logrus.New()

	if display {
		logger.SetOutput(io.MultiWriter(os.Stdout, f))
	} else {
		logger.SetOutput(io.MultiWriter(f))
	}

	logger.SetReportCaller(false)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
