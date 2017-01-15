package model

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		ID        bson.ObjectId
		Email     string
		Password  string
		Token     string
		Followers []string
	}
)
