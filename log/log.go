package log

import (
	"strings"
	"github.com/Sirupsen/logrus"
	"github.com/webus/tanq/utils"
)

var logger *logrus.Logger

// GetLogger - init
func GetLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}
	logger = logrus.New()
	if strings.ToLower(utils.GetEnvVar("debug", "")) == "true" {
		logger.Level = logrus.DebugLevel
	}
	return logger
}
