package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

func GetSession(host string) *mgo.Session {
	log.Println("connect to MongoDB: " + host)
	Session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	return Session
}

