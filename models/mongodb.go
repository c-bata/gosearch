package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

var Session *mgo.Session

func Dialdb(host string) error {
	var err error
	log.Println("connect to MongoDB: localhost")
	Session, err = mgo.Dial(host)
	return err
}

func Closedb() {
	Session.Close()
	log.Println("Close db connection.")
}
