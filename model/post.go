package model

import "gopkg.in/mgo.v2/bson"

type (
	Post struct {
		ID      bson.ObjectId
		To      string
		From    string
		Message string
	}
)
