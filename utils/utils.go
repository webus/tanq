package utils

import (
	"os"
)

// GetEnvVar - get environment variable of get default value
func GetEnvVar(name string, defval string) string {
	data := os.Getenv(name)
	if data == "" {
		return defval
	}
	return data
}
