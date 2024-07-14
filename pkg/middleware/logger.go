package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	logrus "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler, logFilePAth string) http.Handler {
	file, err := os.OpenFile(logFilePAth, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := logrus.New()
	logger.Out = file
	logger.Level = logrus.InfoLevel
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Infof("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		logger.Infof("Completed %s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}
