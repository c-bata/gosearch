package crawler

import (
	"fmt"
	"github.com/c-bata/gosearch/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/c-bata/gosearch/env"
)

func TestGetAllLinks(t *testing.T) {
	assert := assert.New(t)
	var input string = "hoge<a href=\"http://example.com/\">link title</a>hoge\n<a href=\"http://hoge.example.com/\">link title</a>"
	urls := getAllLinks(input)
	assert.Equal("http://example.com/", urls[0])
	assert.Equal("http://hoge.example.com/", urls[1])
}

func DummyCrawledHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hogehoge<a href=\"http://example.com/\">fugafuga")
}

func TestCrawl(t *testing.T) {
	assert := assert.New(t)
	err := models.Dialdb(env.GetDBHost())
	assert.Nil(err)

	ts := httptest.NewServer(http.HandlerFunc(DummyCrawledHandler))
	defer ts.Close()

	msg := make(chan string)
	tocrawl := make(chan URL)
	go crawl(ts.URL, 1, msg, tocrawl)

	for i := 0; i < 2; i++ {
		select {
		case s := <-msg:
			assert.Equal(ts.URL+" is crawled.", s)
		case u := <-tocrawl:
			assert.Equal("http://example.com/", u.url)
			assert.Equal(0, u.depth)
		}
	}
}
