package conf

import (
	"github.com/webus/tanq/utils"
	"github.com/webus/tanq/collections"
)


// GetConf - get current configuration
func GetConf() collections.Configuration {
	cfg := collections.Configuration{}
	cfg.MongoHost = utils.GetEnvVar("MONGO_HOST", "127.0.0.1")
	cfg.MongoDb = utils.GetEnvVar("MONGO_DB", "cdn")
	cfg.MongoCollection = utils.GetEnvVar("MONGO_COLLECTION", "images")
	cfg.MongoGridFS = utils.GetEnvVar("MONGO_GRIDFS", "fs")
	return cfg
}
