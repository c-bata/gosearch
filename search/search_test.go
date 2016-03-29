package search

import (
	"github.com/c-bata/gosearch/env"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/c-bata/gosearch/models"
)

func dropCollection(dbname string) {
	c := models.Session.DB(dbname).C("index")
	c.DropCollection()
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	env.Init()
	err := models.Dialdb(env.GetDBHost())
	assert.Nil(err)
	defer models.Session.Close()
	dropCollection(env.GetDBName())
	c := models.GetIndexCollection(env.GetDBName())

	err = c.Insert(&models.Index{
		Keyword: "word1",
		Url: []string{"http://example.com/"},
	})
	var results []string = Search("word1")
	assert.Equal(1, len(results))
	assert.Equal("http://example.com/", results[0])

	err = c.Insert(&models.Index{
		Keyword: "word2",
		Url: []string{"http://example.com/", "http://hoge.example.com/"},
	})
	results = Search("word2")
	assert.Equal(2, len(results))
}
