package main

import (
	"fmt"
	"github.com/c-bata/gosearch/crawler"
)

func main() {
	msg := make(chan string)
	var seed string = "http://golang.org/"
	go crawler.Crawl(seed, 4, msg)

	for m := range msg {
		fmt.Println(m)
	}
}
