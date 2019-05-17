package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerarMails(t *testing.T) {
	assert := assert.New(t)
	assert.True(Heroe(10, 5))
	assert.False(Heroe(7, 4))
	assert.False(Heroe(4, 5))
	assert.True(Heroe(100, 40))
	assert.False(Heroe(1500, 751))
	assert.False(Heroe(0, 1))
}
