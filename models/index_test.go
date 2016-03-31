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
	c := GetIndexCollection(env.GetDBName())

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
	c := GetIndexCollection(env.GetDBName())
	AddPageToIndex("検索エンジン", "http://example.com")

	var results []Index
	err = c.Find(nil).All(&results)
	assert.Equal(len([]string{"検索", "エンジン"}), len(results))
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	env.Init()
	err := Dialdb(env.GetDBHost())
	assert.Nil(err)
	defer Session.Close()
	dropCollection(env.GetDBName())
	c := GetIndexCollection(env.GetDBName())

	err = c.Insert(&Index{
		Keyword: "word1",
		Url: []string{"http://example.com/"},
	})
	var results []string = Search("word1")
	assert.Equal(1, len(results))
	assert.Equal("http://example.com/", results[0])

	err = c.Insert(&Index{
		Keyword: "word2",
		Url: []string{"http://example.com/", "http://hoge.example.com/"},
	})
	results = Search("word2")
	assert.Equal(2, len(results))
}
