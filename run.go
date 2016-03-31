package main

import (
	"fmt"
	"github.com/c-bata/gosearch/crawler"
	"github.com/c-bata/gosearch/env"
	"github.com/c-bata/gosearch/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/c-bata/gosearch/search"
)

func main() {
	msg := make(chan string)
	var seed string = "http://golang.org/"
	env.Init()
	if err := models.Dialdb(env.GetDBHost()); err != nil {
		fmt.Println("Cannot connect to MongoDB")
		return
	}

	go func() {
		go crawler.Crawl(seed, 4, msg)

		for m := range msg {
			fmt.Println(m)
		}
	} ()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		keyword := c.Query("keyword")
		urls := search.Search(keyword)
		c.JSON(http.StatusOK, gin.H{
			"results": urls,
		})
	})
	router.Run(":8080")
}
