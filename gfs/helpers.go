package gfs

import (
	"strings"
	"path/filepath"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//log "github.com/Sirupsen/logrus"
	"github.com/palantir/stacktrace"
	"github.com/webus/tanq/conf"
	"github.com/webus/tanq/collections"
	logger "github.com/webus/tanq/log"
)

var log = logger.GetLogger()

func (c *MongoGFS) getMongoConnection() {
	log.Debug("Checking connection with Mongo")
	if c.MongoSession != nil && c.MongoDB != nil && c.MongoCollection != nil && c.GFS != nil {
		log.Debug("Connection with Mongo already exists")
		return
	}
	c.Conf = conf.GetConf()
	var err error
	c.MongoSession, err = mgo.Dial(c.Conf.MongoHost)
	if err != nil {
		log.Fatal(stacktrace.Propagate(err, "Error on mgo.Dial"))
	}
	c.MongoSession.SetMode(mgo.Monotonic, true)
	c.MongoDB = c.MongoSession.DB(c.Conf.MongoDb)
	c.MongoCollection = c.MongoDB.C(c.Conf.MongoCollection)
	c.GFS = c.MongoDB.GridFS(c.Conf.MongoGridFS)
}

func (c *MongoGFS) getFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

func (c *MongoGFS) getExistingFileInfoByHash(hash string) *collections.ImageCollection {
	c.getMongoConnection()
	result := collections.ImageCollection{}
	_ = c.MongoCollection.Find(bson.M{"hash":hash}).One(&result)
	if result != (collections.ImageCollection{}) {
		return &result
	}
	return nil
}

func (c *MongoGFS) buildFileName(hash string) string {
	fileInfo := c.getExistingFileInfoByHash(hash)
	if fileInfo == nil {
		// FIXME: ???
		return ""
	}
	fileType := c.getFileExtension(fileInfo.URL)
	return "image" + fileType
}
