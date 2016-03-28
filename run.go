package main

import (
	"fmt"
	"github.com/c-bata/gosearch/crawler"
	"github.com/c-bata/gosearch/env"
	"github.com/c-bata/gosearch/models"
)

func main() {
	msg := make(chan string)
	var seed string = "http://golang.org/"
	env.Init()
	if err := models.Dialdb(env.GetDBHost()); err != nil {
		fmt.Println("Cannot connect to MongoDB")
		return
	}
	go crawler.Crawl(seed, 4, msg)

	for m := range msg {
		fmt.Println(m)
	}
}
