package middleware

import (
	"os"

	logrus "github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func LogInit(logFilePAth string) error {
	file, err := os.OpenFile(logFilePAth, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	Log = logrus.New()
	Log.Out = file
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetReportCaller(true)
	return nil
}
