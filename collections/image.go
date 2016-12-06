package collections

import (
	"gopkg.in/mgo.v2/bson"
)

// ImageCollection - MongoDB collection structure to store files meta info
type ImageCollection struct {
	URL string
	ETag string
	LastModified string
	FileID bson.ObjectId
	ThumbnailFileID bson.ObjectId `bson:",omitempty"`
	Image1FileID bson.ObjectId `bson:",omitempty"`
	Image2FileID bson.ObjectId `bson:",omitempty"`
	Image3FileID bson.ObjectId `bson:",omitempty"`
	Hash string
	Width int `bson:",omitempty"`
	Height int `bson:",omitempty"`
}
