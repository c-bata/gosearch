package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllLinks(t *testing.T) {
	assert := assert.New(t)
	var input string = "hoge<a href=\"http://example.com/\">link title</a>hoge\n<a href=\"http://hoge.example.com/\">link title</a>"
	urls := getAllLinks(input)
	assert.Equal("http://example.com/", urls[0])
	assert.Equal("http://hoge.example.com/", urls[1])
}
