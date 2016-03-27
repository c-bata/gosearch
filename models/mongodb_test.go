package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	assert := assert.New(t)

	session := GetSession("localhost")
	assert.NotNil(session)
	assert.Nil(session.Ping())

	session.Close()
	assert.Panics(func () {session.Ping()}, "Session already closed")
}
