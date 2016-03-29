package search

import (
	"github.com/c-bata/gosearch/models"
	"github.com/c-bata/gosearch/env"
	"gopkg.in/mgo.v2/bson"
)

func Search(keyword string) (urls []string) {
	c := models.GetIndexCollection(env.GetDBName())
	var result models.Index
	c.Find(bson.M{"keyword": keyword}).One(&result)
	return result.Url
}
