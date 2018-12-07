package main

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

const (
	hosts      = "dockercompose_mongodb_1:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "jobs"
)

// NewDB registers a new instance for Mongo
func NewDB() (session *mgo.Session) {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return
}
