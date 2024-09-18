
package common

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func MainLog() *logrus.Logger {
	var logger = logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
