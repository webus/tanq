package gfs

import (
	"github.com/webus/tanq/collections"
	"net/http"
	"github.com/palantir/stacktrace"
	"strings"
	"io"
	"gopkg.in/mgo.v2/bson"
	log "github.com/Sirupsen/logrus"
)

// UploadFileByURL - upload
func (c *MongoGFS) UploadFileByURL(url string) *collections.ImageCollection {
	// 1. upload file by url
	// 2. generate filename extension based on Content-Type
	// 3. make file on GFS
	// 4. copy file to GFS
	// 5. return collections.ImageCollection instance
	log.Debug("???")
	c.getMongoConnection()

	result := collections.ImageCollection{}
	result.Id = bson.NewObjectId()

	log.WithFields(log.Fields{"url": url}).Debug("Upload new image")
	client := http.Client{}
	respGet, err := client.Get(url)
	log.WithFields(log.Fields{"url": url}).Debug("Uploaded image")

	if err != nil {
		log.Fatal(stacktrace.Propagate(err, "Error on client.Get"))
	}
	fileName := "image"
	if strings.ToLower(respGet.Header.Get("Content-Type")) == "image/jpeg" {
		fileName = fileName + ".jpg"
	}
	if strings.ToLower(respGet.Header.Get("Content-Type")) == "image/png" {
		fileName = fileName + ".png"
	}

	file, err := c.GFS.Create(fileName)
	if err != nil {
		log.Fatal(stacktrace.Propagate(err,"Error on gfs.Create"))
	}
	defer file.Close()
	log.WithFields(log.Fields{fileName: fileName}).Debug("Created empty file in GridFS")

	_, err = io.Copy(file, respGet.Body)
	if err != nil {
		log.Fatal(stacktrace.Propagate(err,"Error on Copy"))
	}
	log.WithFields(log.Fields{fileName: fileName}).Debug("File content copied into GridFS file")

	result.URL = url
	result.ETag = respGet.Header.Get("Etag")
	result.LastModified = respGet.Header.Get("Last-Modified")
	result.FileID = file.Id().(bson.ObjectId)
	result.Hash = c.GetImageHashByURL(url)
	err = c.MongoCollection.Insert(&result)
	if err != nil {
		log.Printf("%+v\n", &result)
		log.Fatal(stacktrace.Propagate(err,"Error on insert in MongoDB"))
	}
	log.WithFields(log.Fields{"id": result.Id}).Debug("New ImageCollection created")

	return &result
}
