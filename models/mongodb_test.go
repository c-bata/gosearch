package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	assert := assert.New(t)

	if err := Dialdb(); err != nil {
		t.Error("Failed to connect MongoDB")
	}
	assert.NotNil(Session)
	assert.Nil(Session.Ping())

	Session.Close()
	assert.Panics(func () {Session.Ping()}, "Session already closed")
}
