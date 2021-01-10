package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBlocklistContainsItem(t *testing.T) {
	db := new(UrlinfoDb)
	assert.True(t, db.isBlocklisted("google.com"))
}

func TestBlocklistDoesNotContainItem(t *testing.T) {
	db := new(UrlinfoDb)
	assert.False(t, db.isBlocklisted("hotbot.com"))
}

