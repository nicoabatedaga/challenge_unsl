package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerarMails(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[]string{"rmiranda@ml.com", "mantonieta@ml.com", "rmiranda1@ml.com", "jperez@ml.com"},
		GenerarMails([]string{"Roberto Miranda", "Maria Antonieta", "Raul Miranda", "Jose Perez"}),
		"La lista que se genero no es la que se esperaba")

	assert.Equal(
		[]string{"rmiranda@ml.com", "rmiranda1@ml.com", "rmiranda2@ml.com", "rmiranda3@ml.com"},
		GenerarMails([]string{"Roberto MirAnda", "RAul Miranda", "Raquel MIRANDA", "Rosa Miranda"}),
		"La lista que se genero no es la que se esperaba")
}
