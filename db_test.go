package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBlocklistContainsItem(t *testing.T) {
	assert.True(t, isBlocklisted("google.com"))
}

func TestBlocklistDoesNotContainItem(t *testing.T) {
	assert.False(t, isBlocklisted("hotbot.com"))
}

