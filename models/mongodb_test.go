package models

import (
	"github.com/c-bata/gosearch/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	assert := assert.New(t)
	env.Init()

	err := Dialdb(env.GetDBHost())
	assert.Nil(err)
	assert.NotNil(Session)
	assert.Nil(Session.Ping())

	Session.Close()
	assert.Panics(func() { Session.Ping() }, "Session already closed")
}
