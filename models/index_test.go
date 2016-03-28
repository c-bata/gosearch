package models

import (
	"github.com/c-bata/gosearch/env"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func dropCollection(dbname string) {
	c := Session.DB(dbname).C("index")
	c.DropCollection()
}

func TestAddToIndex(t *testing.T) {
	assert := assert.New(t)
	env.Init()
	err := Dialdb(env.GetDBHost())
	assert.Nil(err)
	defer Session.Close()
	dropCollection(env.GetDBName())
	c := getIndexCollection(env.GetDBName())

	// When add a new word
	addToIndex("keyword1", "http://example.com")
	result := &Index{}
	err = c.Find(bson.M{"keyword": "keyword1"}).One(&result)
	assert.Equal(1, len(result.Url))

	// When add a different url
	addToIndex("keyword1", "http://hoge.example.com")
	err = c.Find(bson.M{"keyword": "keyword1"}).One(&result)
	assert.Equal(2, len(result.Url))

	// When add a same url
	addToIndex("keyword1", "http://hoge.example.com")
	err = c.Find(bson.M{"keyword": "keyword1"}).One(&result)
	assert.Equal(2, len(result.Url))

	// When add a new word
	addToIndex("keyword2", "http://fuga.example.com")

	var results []Index
	err = c.Find(nil).All(&results)
	assert.Equal(2, len(results))
}

func TestAddPageToIndex(t *testing.T) {
	assert := assert.New(t)
	env.Init()
	err := Dialdb(env.GetDBHost())
	assert.Nil(err)
	defer Session.Close()
	dropCollection(env.GetDBName())
	c := getIndexCollection(env.GetDBName())
	AddPageToIndex("検索エンジン", "http://example.com")

	var results []Index
	err = c.Find(nil).All(&results)
	assert.Equal(len([]string{"検索", "エンジン"}), len(results))
}
