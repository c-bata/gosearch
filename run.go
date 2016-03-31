package main

import (
	"fmt"
	"github.com/c-bata/gosearch/crawler"
	"github.com/c-bata/gosearch/env"
	"github.com/c-bata/gosearch/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	resp := make(chan crawler.CrawlResponse)
	var seed string = "http://golang.org/"
	env.Init()
	if err := models.Dialdb(env.GetDBHost()); err != nil {
		fmt.Println("Cannot connect to MongoDB")
		return
	}

	go func() {
		go crawler.Crawl(seed, 4, resp)

		for r := range resp {
			// TODO: Add add page to index test
			// TODO: Remove tags
			// TODO: Skip binary and static files(js, css, img)
			//       Check Content-Type (json, html, xml, etc)
			log.Printf("%d : %s", r.StatusCode, r.Url)
			models.AddPageToIndex(string(r.Body), r.Url)
		}
	} ()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		keyword := c.Query("keyword")
		urls := models.Search(keyword)
		c.JSON(http.StatusOK, gin.H{
			"results": urls,
		})
	})
	router.Run(":8080")
}
