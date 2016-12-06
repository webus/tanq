package utils

import (
	"os"
	"strings"
)

// GetEnvVar - get environment variable of get default value
func GetEnvVar(name string, defval string) string {
	data := os.Getenv(name)
	dataUpper := os.Getenv(strings.ToUpper(name))
	if data == "" && dataUpper == "" {
		return defval
	} else if data != "" {
		return data
	} else if dataUpper != "" {
		return dataUpper
	}
	return defval
}
