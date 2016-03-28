package models

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func dropCollection(assert *assert.Assertions, db string, col string) {
	c := Session.DB("gosearch").C("index")
	c.DropCollection()
}

func TestAddToIndex(t *testing.T) {
	assert := assert.New(t)
	err := Dialdb()
	assert.Nil(err)
	defer Session.Close()
	dropCollection(assert, "gosearch", "index")
	c := Session.DB("gosearch").C("index")

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
	err := Dialdb()
	assert.Nil(err)
	defer Session.Close()
	dropCollection(assert, "gosearch", "index")
	c := Session.DB("gosearch").C("index")
	AddPageToIndex("検索エンジン", "http://example.com")

	var results []Index
	err = c.Find(nil).All(&results)
	assert.Equal(len([]string{"検索", "エンジン"}), len(results))
}
