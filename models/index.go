package models

import (
	"github.com/c-bata/gosearch/env"
	"github.com/ikawaha/kagome/tokenizer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Index struct {
	Keyword string   `bson:"keyword"`
	Url     []string `bson:"url"`
}

func GetIndexCollection(db string) *mgo.Collection {
	return Session.DB(db).C("index")
}

func contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func addToIndex(keyword string, url string) (err error) {
	c := GetIndexCollection(env.GetDBName())

	result := &Index{}
	if err := c.Find(bson.M{"keyword": keyword}).One(&result); err != nil {
		err = c.Insert(&Index{Keyword: keyword, Url: []string{url}})
	} else if !contains(url, result.Url) {
		err = c.Update(bson.M{"keyword": keyword}, bson.M{"$push": bson.M{"url": url}})
	}
	return
}

func AddPageToIndex(body string, url string) {
	t := tokenizer.New()
	tokens := t.Tokenize(body)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		addToIndex(token.Surface, url)
	}
}

func Search(keyword string) (urls []string) {
	c := GetIndexCollection(env.GetDBName())
	var result Index
	c.Find(bson.M{"keyword": keyword}).One(&result)
	return result.Url
}
