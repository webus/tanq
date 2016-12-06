package gfs

import (
	"io"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/webus/tanq/collections"
	"github.com/palantir/stacktrace"
)

// MongoGFS - mongo gfs
type MongoGFS struct {
	MongoSession *mgo.Session
	MongoDB *mgo.Database
	MongoCollection *mgo.Collection
	GFS *mgo.GridFS
	Conf collections.Configuration
}

// Close - close
func (c *MongoGFS) Close() {
	c.MongoSession.Close()
}

// GetFileByHash - hash
func (c *MongoGFS) GetFileByHash(hash string) (*mgo.GridFile, error) {
	c.getMongoConnection()
	result := collections.ImageCollection{}
	_ = c.MongoCollection.Find(bson.M{"hash":hash}).One(&result)
	if result != (collections.ImageCollection{}) {
		file, err := c.GFS.OpenId(result.FileID)
		if err != nil {
			return nil, stacktrace.Propagate(err,"Error on gfs.OpenId")
		}
		return file, nil
	}
	return  nil, nil
}

// GetFileByHashFull - full. return file bytes
func (c *MongoGFS) GetFileByHashFull(hash string) ([]byte, error) {
	c.getMongoConnection()
	gridFile, err := c.GetFileByHash(hash)
	if err != nil {
		return nil, stacktrace.Propagate(err,"Error on gfs.OpenId")
	}
	if gridFile != nil {
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, gridFile)
		return buf.Bytes(), nil
	}
	return nil, nil
}

// GetImageHashByURLFromDB - url
func (c *MongoGFS) GetImageHashByURLFromDB(url string) string {
	c.getMongoConnection()
	result := collections.ImageCollection{}
	_ = c.MongoCollection.Find(bson.M{"url":url}).One(&result)
	if result != (collections.ImageCollection{}) {
		return result.Hash
	}
	return ""
}

// GetImageHashByURL - get
func (c *MongoGFS) GetImageHashByURL(url string) string {
	hasher := sha1.New()
	return base64.URLEncoding.EncodeToString(hasher.Sum([]byte(url)))
}
